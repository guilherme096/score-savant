USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

-- Alter the UDF to get player information with pagination, sorting, and search filters
ALTER FUNCTION dbo.GetPlayersWithPagination
(
    @PageNumber INT,
    @PageSize INT,
    @OrderBy NVARCHAR(50) = NULL,
    @OrderDirection NVARCHAR(4) = NULL,
    @SearchPlayerName NVARCHAR(255) = NULL,
    @SearchClubName NVARCHAR(255) = NULL,
    @SearchPositionName NVARCHAR(255) = NULL,
    @SearchNationName NVARCHAR(255) = NULL,
    @SearchLeagueName NVARCHAR(255) = NULL,
    @MinWage DECIMAL(18,2) = NULL,
    @MaxWage DECIMAL(18,2) = NULL,
    @MinValue DECIMAL(18,2) = NULL,
    @MaxValue DECIMAL(18,2) = NULL,
    @MinAge INT = NULL,
    @MaxAge INT = NULL,
    @MinReleaseClause DECIMAL(18,2) = NULL,
    @MaxReleaseClause DECIMAL(18,2) = NULL
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
                    CASE WHEN @OrderBy = 'PlayerName' AND @OrderDirection = 'ASC' THEN p.name END ASC,
                    CASE WHEN @OrderBy = 'PlayerName' AND @OrderDirection = 'DESC' THEN p.name END DESC,
                    CASE WHEN @OrderBy = 'Age' AND @OrderDirection = 'ASC' THEN p.age END ASC,
                    CASE WHEN @OrderBy = 'Age' AND @OrderDirection = 'DESC' THEN p.age END DESC,
                    CASE WHEN @OrderBy = 'Position' AND @OrderDirection = 'ASC' THEN pos.name END ASC,
                    CASE WHEN @OrderBy = 'Position' AND @OrderDirection = 'DESC' THEN pos.name END DESC,
                    CASE WHEN @OrderBy = 'Club' AND @OrderDirection = 'ASC' THEN cl.name END ASC,
                    CASE WHEN @OrderBy = 'Club' AND @OrderDirection = 'DESC' THEN cl.name END DESC,
                    CASE WHEN @OrderBy = 'Wage' AND @OrderDirection = 'ASC' THEN c.wage END ASC,
                    CASE WHEN @OrderBy = 'Wage' AND @OrderDirection = 'DESC' THEN c.wage END DESC,
                    CASE WHEN @OrderBy = 'Value' AND @OrderDirection = 'ASC' THEN p.value END ASC,
                    CASE WHEN @OrderBy = 'Value' AND @OrderDirection = 'DESC' THEN p.value END DESC,
                    CASE WHEN @OrderBy = 'Nation' AND @OrderDirection = 'ASC' THEN n.name END ASC,
                    CASE WHEN @OrderBy = 'Nation' AND @OrderDirection = 'DESC' THEN n.name END DESC,
                    CASE WHEN @OrderBy = 'League' AND @OrderDirection = 'ASC' THEN l.name END ASC,
                    CASE WHEN @OrderBy = 'League' AND @OrderDirection = 'DESC' THEN l.name END DESC,
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
        INNER JOIN
            Nation n ON p.nation_id = n.nation_id
        INNER JOIN
            League l ON cl.league_id = l.league_id
        WHERE
            (@SearchPlayerName IS NULL OR p.name LIKE '%' + @SearchPlayerName + '%')
            AND (@SearchClubName IS NULL OR cl.name LIKE '%' + @SearchClubName + '%')
            AND (@SearchPositionName IS NULL OR pos.name LIKE '%' + @SearchPositionName + '%')
            AND (@SearchNationName IS NULL OR n.name LIKE '%' + @SearchNationName + '%')
            AND (@SearchLeagueName IS NULL OR l.name LIKE '%' + @SearchLeagueName + '%')
            AND (@MinWage IS NULL OR c.wage >= @MinWage)
            AND (@MaxWage IS NULL OR c.wage <= @MaxWage)
            AND (@MinValue IS NULL OR p.value >= @MinValue)
            AND (@MaxValue IS NULL OR p.value <= @MaxValue)
            AND (@MinAge IS NULL OR p.age >= @MinAge)
            AND (@MaxAge IS NULL OR p.age <= @MaxAge)
            AND (@MinReleaseClause IS NULL OR c.release_clause >= @MinReleaseClause)
            AND (@MaxReleaseClause IS NULL OR c.release_clause <= @MaxReleaseClause)
    ) AS PlayersWithRowNum
    WHERE
        RowNum > (@PageNumber - 1) * @PageSize
        AND RowNum <= @PageNumber * @PageSize
)
GO
       
