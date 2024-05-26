USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

-- Create the UDF to get role name and role ID by player ID
CREATE FUNCTION dbo.GetRoleByPlayerID
(
    @PlayerID INT
)
RETURNS TABLE
AS
RETURN
(
    SELECT
        r.role_id AS RoleID,
        r.name AS RoleName
    FROM
        Player p
    INNER JOIN
        PlayerRole playrole ON playrole.player_id = p.player_id
    INNER JOIN
        Role r ON r.role_id = playrole.role_position_id
    WHERE
        p.player_id = @PlayerID
)
GO
