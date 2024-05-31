USE [p5g5]
GO

CREATE TRIGGER trg_after_delete_club
ON Club
INSTEAD OF DELETE
AS
BEGIN
    DECLARE @club_id INT;

    -- Loop through each club to be deleted
    DECLARE deleted_club_cursor CURSOR FOR
    SELECT club_id FROM DELETED;

    OPEN deleted_club_cursor;

    FETCH NEXT FROM deleted_club_cursor INTO @club_id;

    WHILE @@FETCH_STATUS = 0
    BEGIN
        -- Delete associated players first
        DELETE FROM Player WHERE club_id = @club_id;

        -- Finally, delete the club record
        DELETE FROM Club WHERE club_id = @club_id;

        FETCH NEXT FROM deleted_club_cursor INTO @club_id;
    END

    CLOSE deleted_club_cursor;
    DEALLOCATE deleted_club_cursor;
END;
GO
