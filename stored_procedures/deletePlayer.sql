USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE dbo.DeletePlayer
    @player_id INT
AS
BEGIN
    SET NOCOUNT ON;

    -- Delete player from related tables
    DELETE FROM Contract WHERE player_id = @player_id;

    -- Get player role and position IDs
    DECLARE @role_position_id INT;
    DECLARE @position_id INT;
    DECLARE @position NVARCHAR(255);
    DECLARE @player_type INT;

    SELECT @role_position_id = role_position_id FROM PlayerRole WHERE player_id = @player_id;
    DELETE FROM PlayerRole WHERE player_id = @player_id;

    SELECT @position_id = position_id FROM PlayerPosition WHERE player_id = @player_id;
    DELETE FROM PlayerPosition WHERE player_id = @player_id;

    SELECT @position = name FROM Position WHERE position_id = @position_id;

    -- Determine player type based on position
    IF @position = 'GK'
    BEGIN
        SET @player_type = 1;
    END
    ELSE
    BEGIN
        SET @player_type = 0;
    END

    -- Delete from specific player type tables
    IF @player_type = 0
    BEGIN
        DELETE FROM Outfield_Player WHERE player_id = @player_id;
        DELETE FROM OutfieldAttributeRating WHERE player_id = @player_id;
    END
    ELSE
    BEGIN
        DELETE FROM Goalkeeper WHERE player_id = @player_id;
        DELETE FROM GoalkeeperAttributeRating WHERE player_id = @player_id;
    END

    -- Finally, delete the player
    DELETE FROM Player WHERE player_id = @player_id;
END
GO
