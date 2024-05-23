use [p5g5]
go

set ansi_nulls on
go

set quoted_identifier on
go

create procedure [dbo].[AddPlayerRole]
    @role int,
    @player int
as
begin
    set nocount on;

    declare @role_id int;
    declare @player_id int;

    -- Get Role ID
    select @role_id = role_id from Role where role_id = @role;
    if @role_id is null
    begin
        raiserror('Role not found: %d', 16, 1, @role);
        return;
    end

    -- Get or insert player
    select @player_id = player_id from Player where player_id = @player;
    if @player_id is null
    begin
        raiserror('Player not found: %d', 16, 1, @player);
        return;
    end

    -- Get or insert player role
    if not exists (select * from PlayerRole where player_id = @player_id and role_id = @role_id)
    begin
        insert into PlayerRole (player_id, role_id) values (@player_id, @role_id);
    end
end