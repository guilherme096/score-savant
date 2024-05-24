USE [p5g5]
GO

CREATE TRIGGER trg_RemoveStarredPlayer
ON Player
AFTER DELETE
AS
BEGIN
    DELETE FROM StaredPlayers
    WHERE player_id IN (SELECT player_id FROM deleted);
END;
GO
