USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

-- Create the UDF to get club information with pagination, sorting, and search filters
ALTER FUNCTION dbo.GetClubsWithPagination
(
    @PageNumber INT,
    @PageSize INT,
    @OrderBy NVARCHAR(50) = NULL,
    @OrderDirection NVARCHAR(4) = NULL,
    @SearchClubName NVARCHAR(255) = NULL,
    @SearchLeagueName NVARCHAR(255) = NULL,
    @SearchNationName NVARCHAR(255) = NULL,
    @MinPlayerCount INT = NULL,
    @MaxPlayerCount INT = NULL,
    @MinWageTotal DECIMAL(18,2) = NULL,
    @MaxWageTotal DECIMAL(18,2) = NULL,
    @MinValueTotal DECIMAL(18,2) = NULL,
    @MaxValueTotal DECIMAL(18,2) = NULL
)
RETURNS TABLE
AS
RETURN
(
    SELECT 
        ClubID,
        ClubName,
        Nation,
        League,
        PlayerCount,
        WageTotal,
        ValueTotal
    FROM
    (
        SELECT
            c.club_id AS ClubID,
            c.name AS ClubName,
            n.name AS Nation,
            l.name AS League,
            c.player_count AS PlayerCount,
            c.wage_total AS WageTotal,
            c.value_total AS ValueTotal,
            ROW_NUMBER() OVER (
                ORDER BY
                    CASE WHEN @OrderBy = 'ClubName' AND @OrderDirection = 'ASC' THEN c.name END ASC,
                    CASE WHEN @OrderBy = 'ClubName' AND @OrderDirection = 'DESC' THEN c.name END DESC,
                    CASE WHEN @OrderBy = 'Nation' AND @OrderDirection = 'ASC' THEN n.name END ASC,
                    CASE WHEN @OrderBy = 'Nation' AND @OrderDirection = 'DESC' THEN n.name END DESC,
                    CASE WHEN @OrderBy = 'League' AND @OrderDirection = 'ASC' THEN l.name END ASC,
                    CASE WHEN @OrderBy = 'League' AND @OrderDirection = 'DESC' THEN l.name END DESC,
                    CASE WHEN @OrderBy = 'PlayerCount' AND @OrderDirection = 'ASC' THEN c.player_count END ASC,
                    CASE WHEN @OrderBy = 'PlayerCount' AND @OrderDirection = 'DESC' THEN c.player_count END DESC,
                    CASE WHEN @OrderBy = 'WageTotal' AND @OrderDirection = 'ASC' THEN c.wage_total END ASC,
                    CASE WHEN @OrderBy = 'WageTotal' AND @OrderDirection = 'DESC' THEN c.wage_total END DESC,
                    CASE WHEN @OrderBy = 'ValueTotal' AND @OrderDirection = 'ASC' THEN c.value_total END ASC,
                    CASE WHEN @OrderBy = 'ValueTotal' AND @OrderDirection = 'DESC' THEN c.value_total END DESC,
                    -- Default order
                    c.club_id
            ) AS RowNum
        FROM
            Club c
        INNER JOIN
            Nation n ON c.nation_id = n.nation_id
        INNER JOIN
            League l ON c.league_id = l.league_id
        WHERE
            (@SearchClubName IS NULL OR c.name LIKE '%' + @SearchClubName + '%')
            AND (@SearchLeagueName IS NULL OR l.name LIKE '%' + @SearchLeagueName + '%')
            AND (@SearchNationName IS NULL OR n.name LIKE '%' + @SearchNationName + '%')
            AND (@MinPlayerCount IS NULL OR c.player_count >= @MinPlayerCount)
            AND (@MaxPlayerCount IS NULL OR c.player_count <= @MaxPlayerCount)
            AND (@MinWageTotal IS NULL OR c.wage_total >= @MinWageTotal)
            AND (@MaxWageTotal IS NULL OR c.wage_total <= @MaxWageTotal)
            AND (@MinValueTotal IS NULL OR c.value_total >= @MinValueTotal)
            AND (@MaxValueTotal IS NULL OR c.value_total <= @MaxValueTotal)
    ) AS ClubsWithRowNum
    WHERE
        RowNum > (@PageNumber - 1) * @PageSize
        AND RowNum <= @PageNumber * @PageSize
)
GO
