USE [p5g5]
GO

CREATE TRIGGER trg_after_delete_nation
ON Nation
AFTER DELETE
AS
BEGIN
    -- Declare variables
    DECLARE @nation_id INT;

    -- Loop through each deleted nation
    DECLARE deleted_nation_cursor CURSOR FOR
    SELECT nation_id FROM DELETED;

    OPEN deleted_nation_cursor;

    FETCH NEXT FROM deleted_nation_cursor INTO @nation_id;

    WHILE @@FETCH_STATUS = 0
    BEGIN
        -- Delete leagues associated with the deleted nation
        DELETE FROM League WHERE nation_id = @nation_id;

        -- Delete clubs associated with the deleted nation
        DELETE FROM Club WHERE nation_id = @nation_id;

        -- Delete players of the deleted nationality
        DELETE FROM Player WHERE nation_id = @nation_id;

        FETCH NEXT FROM deleted_nation_cursor INTO @nation_id;
    END

    CLOSE deleted_nation_cursor;
    DEALLOCATE deleted_nation_cursor;
END;
GO
