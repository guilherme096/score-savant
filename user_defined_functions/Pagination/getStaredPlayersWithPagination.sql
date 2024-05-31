USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE FUNCTION dbo.GetStaredPlayersWithPagination
(
    @PageNumber INT,
    @PageSize INT,
    @OrderBy NVARCHAR(50) = NULL,
    @OrderDirection NVARCHAR(4) = NULL
)
RETURNS TABLE
AS
RETURN
(
    SELECT 
        PlayerID,
        PlayerName,
        Position,
        Club,
        Wage,
        Value,
        Nation,
        League,
        Age,
        ReleaseClause
    FROM
    (
        SELECT
            p.player_id AS PlayerID,
            p.name AS PlayerName,
            pos.name AS Position,
            cl.name AS Club,
            c.wage AS Wage,
            p.value AS Value,
            n.name AS Nation,
            l.name AS League,
            p.age AS Age,
            c.release_clause AS ReleaseClause,
            ROW_NUMBER() OVER (
                ORDER BY
                    p.player_id
            ) AS RowNum
        FROM
            StaredPlayers sp
        INNER JOIN
            Player p ON sp.player_id = p.player_id
        INNER JOIN
            PlayerPosition playpos ON playpos.player_id = p.player_id
        INNER JOIN
            Position pos ON pos.position_id = playpos.position_id
        INNER JOIN
            Club cl ON p.club_id = cl.club_id
        INNER JOIN
            Contract c ON p.player_id = c.player_id
        INNER JOIN
            Nation n ON p.nation_id = n.nation_id
        INNER JOIN
            League l ON cl.league_id = l.league_id
    ) AS StaredPlayersWithRowNum
    WHERE
        RowNum > (@PageNumber - 1) * @PageSize
        AND RowNum <= @PageNumber * @PageSize
)
GO
