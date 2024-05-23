USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

-- Create the UDF to get position name and position ID by player ID
CREATE FUNCTION dbo.GetPositionByPlayerID
(
    @PlayerID INT
)
RETURNS TABLE
AS
RETURN
(
    SELECT
        pos.position_id AS PositionID,
        pos.name AS PositionName
    FROM
        Player p
    INNER JOIN
        PlayerPosition playpos ON playpos.player_id = p.player_id
    INNER JOIN
        Position pos ON pos.position_id = playpos.position_id
    WHERE
        p.player_id = @PlayerID
)
GO
