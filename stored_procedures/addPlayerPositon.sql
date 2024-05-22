use [p5g5]
go

set ansi_nulls on
go

set quoted_identifier on
go

create procedure [dbo].[AddPlayerPosition]
    @position int,
    @player int
as
begin
    set nocount on;

    declare @position_id int;
    declare @player_id int;

    -- Get Position ID
    select @position_id = position_id from Position where position_id = @position;
    if @position_id is null
    begin
        raiserror('Position not found: %d', 16, 1, @position);
        return;
    end

    -- Get or insert player
    select @player_id = player_id from Player where player_id = @player;
    if @player_id is null
    begin
        raiserror('Player not found: %d', 16, 1, @player);
        return;
    end

    -- Get or insert player position
    if not exists (select * from PlayerPosition where player_id = @player_id and position_id = @position_id)
    begin
        insert into PlayerPosition (player_id, position_id) values (@player_id, @position_id);
    end
end