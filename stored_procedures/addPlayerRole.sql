USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[AddPlayerRole]
    @role INT,
    @player INT
AS
BEGIN
    SET NOCOUNT ON;

    DECLARE @role_id INT;
    DECLARE @player_id INT;

    -- Get Role ID
    SELECT @role_id = role_id FROM Role WHERE role_id = @role;
    IF @role_id IS NULL
    BEGIN
        RAISERROR('Role not found: %d', 16, 1, @role);
        RETURN;
    END

    -- Get Player ID
    SELECT @player_id = player_id FROM Player WHERE player_id = @player;
    IF @player_id IS NULL
    BEGIN
        RAISERROR('Player not found: %d', 16, 1, @player);
        RETURN;
    END

    -- Insert Player Role if not exists
    IF NOT EXISTS (SELECT * FROM PlayerRole WHERE player_id = @player_id AND role_position_id = @role_id)
    BEGIN
        INSERT INTO PlayerRole (player_id, role_position_id) VALUES (@player_id, @role_id);
    END
END
