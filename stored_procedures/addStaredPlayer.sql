USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE dbo.AddStaredPlayer
(
    @PlayerID INT
)
AS
BEGIN
    SET NOCOUNT ON;

    -- Verifica se o jogador já está na lista de favoritos
    IF EXISTS (SELECT 1 FROM StaredPlayers WHERE player_id = @PlayerID)
    BEGIN
        PRINT 'Player is already in the stared players list.';
        RETURN;
    END

    -- Insere o jogador na lista de favoritos
    INSERT INTO StaredPlayers (player_id)
    VALUES (@PlayerID);

    PRINT 'Player successfully added to the stared players list.';
END;
GO
