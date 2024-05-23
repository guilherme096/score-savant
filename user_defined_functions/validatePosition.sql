USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE FUNCTION [dbo].[ValidatePosition]
(
    @position NVARCHAR(255)
)
RETURNS INT
AS
BEGIN
    DECLARE @position_id INT;

    -- Validate position
    SELECT @position_id = position_id FROM Position WHERE name = @position;

    -- Return position_id (NULL if not found)
    RETURN @position_id;
END
GO
