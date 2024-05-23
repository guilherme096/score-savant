USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

-- Create the UDF to get distinct roles by position ID
CREATE FUNCTION dbo.GetRolesByPositionID
(
    @position_id INT
)
RETURNS TABLE
AS
RETURN
(
    SELECT DISTINCT
        r.role_id AS RoleID,
        r.name AS RoleName
    FROM
        RolePosition rp
    INNER JOIN
        ROLE r ON rp.role_position = r.role_id
    WHERE
        rp.position_id = @position_id
)
GO
