use [p5g5]
go

set ansi_nulls on
go

set quoted_identifier on
go

create procedure [dbo].[ValidateRole]
    @role nvarchar(255)
as
begin
    set nocount on;

    declare @role_id int;

    -- Validate role
    select @role_id = role_id from PlayerRole where name = @role;
    if @role_id is null
    begin
        raiserror('Role not found: %s', 16, 1, @role);
        return;
    end
end