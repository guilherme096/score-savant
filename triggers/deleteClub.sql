USE [p5g5]
GO

CREATE TRIGGER trg_after_delete_club
ON Club
AFTER DELETE
AS
BEGIN
    DECLARE @club_id INT;

    -- Get the club_id from the deleted row
    SELECT @club_id = club_id
    FROM DELETED;

    -- Delete players associated with the deleted club
    DELETE FROM Player
    WHERE club_id = @club_id;
END;
GO
