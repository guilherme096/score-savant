use [p5g5]
go

set ansi_nulls on
go

set quoted_identifier on
go

create procedure [dbo].[AddLeague]
    @league nvarchar(255),
    @nation int
as
begin
    set nocount on;

    declare @nation_id int;
    declare @league_id int;

    -- Get Nation ID
    select @nation_id = nation_id from Nation where nation_id = @nation;
    if @nation_id is null
    begin
        raiserror('Nation not found: %d', 16, 1, @nation);
        return;
    end

    -- Get or insert league
    select @league_id = league_id from League where name = @league;
    if @league_id is null
    begin
        insert into League (name, nation_id) values (@league, @nation_id);
        set @league_id = scope_identity();
    end
end
