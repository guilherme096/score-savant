USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE FUNCTION [dbo].[ValidateRole]
(
    @role NVARCHAR(255)
)
RETURNS INT
AS
BEGIN
    DECLARE @role_id INT;

    -- Validate role
    SELECT @role_id = role_id FROM Role WHERE name = @role;

    -- Return role_id (NULL if not found)
    RETURN @role_id;
END
GO
