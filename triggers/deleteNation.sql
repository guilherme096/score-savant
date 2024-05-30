USE [p5g5]
GO

CREATE TRIGGER trg_after_delete_nation
ON Nation
AFTER DELETE
AS
BEGIN
    DECLARE @nation_id INT;

    -- Get the nation_id from the deleted row
    SELECT @nation_id = nation_id
    FROM DELETED;

    -- Delete players of the deleted nationality
    DELETE FROM Player
    WHERE nation_id = @nation_id;

    -- Delete leagues associated with the deleted nation
    DELETE FROM League
    WHERE nation_id = @nation_id;
END;
GO
