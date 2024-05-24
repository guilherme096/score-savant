USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [dbo].[GetTechnicalAtt]
AS
SELECT
    ta.att_id AS AttributeID
FROM
    TechnicalAttribute ta
GO
