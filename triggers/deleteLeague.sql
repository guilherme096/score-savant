USE [p5g5]
GO

CREATE TRIGGER trg_after_delete_league
ON League
AFTER DELETE
AS
BEGIN
    DECLARE @league_id INT;

    -- Get the league_id from the deleted row
    SELECT @league_id = league_id
    FROM DELETED;

    -- Delete clubs associated with the deleted league
    DELETE FROM Club
    WHERE league_id = @league_id;
END;
GO
