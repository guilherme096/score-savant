USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [dbo].[GetGoalkeepingAtt]
AS
SELECT
    ga.att_id AS AttributeID
FROM
    Goalkeeping_Att ga
GO
