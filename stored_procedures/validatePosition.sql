use [p5g5]
go

set ansi_nulls on
go

set quoted_identifier on
go

create procedure [dbo].[ValidatePosition]
    @position nvarchar(255)
as
begin
    set nocount on;

    declare @position_id int;

    -- Validate position
    select @position_id = position_id from Position where name = @position;
    if @position_id is null
    begin
        raiserror('Position not found: %s', 16, 1, @position);
        return;
    end
end