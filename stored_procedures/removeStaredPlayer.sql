USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE dbo.RemoveStaredPlayer
(
    @PlayerID INT
)
AS
BEGIN
    SET NOCOUNT ON;

    -- Verifica se o jogador está na lista de favoritos
    IF NOT EXISTS (SELECT 1 FROM StaredPlayers WHERE player_id = @PlayerID)
    BEGIN
        PRINT 'Player is not in the stared players list.';
        RETURN;
    END

    -- Remove o jogador da lista de favoritos
    DELETE FROM StaredPlayers
    WHERE player_id = @PlayerID;

    PRINT 'Player successfully removed from the stared players list.';
END;
GO
