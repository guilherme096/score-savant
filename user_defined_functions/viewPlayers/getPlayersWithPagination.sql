USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

-- Create the UDF to get player information with pagination and predefined sorting options
CREATE FUNCTION dbo.GetPlayersWithPagination
(
    @PageNumber INT,
    @PageSize INT,
    @OrderBy NVARCHAR(50),
    @OrderDirection NVARCHAR(4)
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
        Value
    FROM
    (
        SELECT
            p.player_id AS PlayerID,
            p.name AS PlayerName,
            pos.name AS Position,
            cl.name AS Club,
            c.wage AS Wage,
            p.value AS Value,
            ROW_NUMBER() OVER (
                ORDER BY
                    CASE WHEN @OrderBy = 'PlayerName' AND @OrderDirection = 'ASC' THEN p.name END ASC,
                    CASE WHEN @OrderBy = 'PlayerName' AND @OrderDirection = 'DESC' THEN p.name END DESC,
                    CASE WHEN @OrderBy = 'Position' AND @OrderDirection = 'ASC' THEN pos.position_id END ASC,
                    CASE WHEN @OrderBy = 'Position' AND @OrderDirection = 'DESC' THEN pos.position_id END DESC,
                    CASE WHEN @OrderBy = 'Club' AND @OrderDirection = 'ASC' THEN cl.name END ASC,
                    CASE WHEN @OrderBy = 'Club' AND @OrderDirection = 'DESC' THEN cl.name END DESC,
                    CASE WHEN @OrderBy = 'Wage' AND @OrderDirection = 'ASC' THEN c.wage END ASC,
                    CASE WHEN @OrderBy = 'Wage' AND @OrderDirection = 'DESC' THEN c.wage END DESC,
                    CASE WHEN @OrderBy = 'Value' AND @OrderDirection = 'ASC' THEN p.value END ASC,
                    CASE WHEN @OrderBy = 'Value' AND @OrderDirection = 'DESC' THEN p.value END DESC,
                    -- Default order
                    p.player_id
            ) AS RowNum
        FROM
            Player p
        INNER JOIN
            PlayerPosition playpos ON playpos.player_id = p.player_id
        INNER JOIN
            Position pos ON pos.position_id = playpos.position_id
        INNER JOIN
            Club cl ON p.club_id = cl.club_id
        INNER JOIN
            Contract c ON p.player_id = c.player_id
    ) AS Paged
    WHERE Paged.RowNum BETWEEN (@PageNumber - 1) * @PageSize + 1 AND @PageNumber * @PageSize
)
GO
