USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

-- Create the UDF to get key attributes by role ID
CREATE FUNCTION dbo.GetKeyAttributesByRoleID
(
    @role_id INT
)
RETURNS TABLE
AS
RETURN
(
    SELECT
        ka.role_id AS RoleID,
        ka.attribute_id AS KeyAttributeID
    FROM
        KeyAttributes ka
    WHERE
        ka.role_id = @role_id
)
GO
