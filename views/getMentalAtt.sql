USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [dbo].[GetMentalAtt]
AS
SELECT
    ma.att_id AS AttributeID
FROM
    Mental_Att ma
GO
