USE [p5g5]
GO

CREATE TRIGGER trg_after_delete_players
ON Player
INSTEAD OF DELETE
AS
BEGIN
    DECLARE @player_id INT;

    -- Loop through each player to be deleted
    DECLARE deleted_player_cursor CURSOR FOR
    SELECT player_id FROM DELETED;

    OPEN deleted_player_cursor;

    FETCH NEXT FROM deleted_player_cursor INTO @player_id;

    WHILE @@FETCH_STATUS = 0
    BEGIN
        -- Delete associated entries from dependent tables
        DELETE FROM Contract WHERE player_id = @player_id;
        DELETE FROM PlayerRole WHERE player_id = @player_id;
        DELETE FROM PlayerPosition WHERE player_id = @player_id;
        DELETE FROM Outfield_Player WHERE player_id = @player_id;
        DELETE FROM Goalkeeper WHERE player_id = @player_id;
        DELETE FROM OutfieldAttributeRating WHERE player_id = @player_id;
        DELETE FROM GoalkeeperAttributeRating WHERE player_id = @player_id;
        DELETE FROM StaredPlayers WHERE player_id = @player_id;

        -- Finally, delete the player record
        DELETE FROM Player WHERE player_id = @player_id;

        FETCH NEXT FROM deleted_player_cursor INTO @player_id;
    END

    CLOSE deleted_player_cursor;
    DEALLOCATE deleted_player_cursor;
END;
GO
