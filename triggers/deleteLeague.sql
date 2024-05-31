USE [p5g5]
GO

CREATE TRIGGER trg_after_delete_league
ON League
INSTEAD OF DELETE
AS
BEGIN
    DECLARE @league_id INT;

    -- Loop through each league to be deleted
    DECLARE deleted_league_cursor CURSOR FOR
    SELECT league_id FROM DELETED;

    OPEN deleted_league_cursor;

    FETCH NEXT FROM deleted_league_cursor INTO @league_id;

    WHILE @@FETCH_STATUS = 0
    BEGIN
        -- Delete associated clubs (and their players) first
        DELETE FROM Club WHERE league_id = @league_id;

        -- Finally, delete the league record
        DELETE FROM League WHERE league_id = @league_id;

        FETCH NEXT FROM deleted_league_cursor INTO @league_id;
    END

    CLOSE deleted_league_cursor;
    DEALLOCATE deleted_league_cursor;
END;
GO
