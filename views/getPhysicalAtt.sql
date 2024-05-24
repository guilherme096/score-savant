USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [dbo].[GetPhysicalAtt]
AS
SELECT
    pa.att_id AS AttributeID
FROM
    Physical_Att pa
GO
