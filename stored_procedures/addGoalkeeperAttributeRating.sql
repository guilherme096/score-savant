USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

ALTER PROCEDURE dbo.AddGoalkeeperAttributeRating
(
    @PlayerID INT,
    @Attributes NVARCHAR(MAX)
)
AS
BEGIN
    SET NOCOUNT ON;
    BEGIN TRY
        BEGIN TRANSACTION;

        -- Verify and return attributes
        DECLARE @VerifiedAttributes TABLE (Attribute NVARCHAR(255), Rating INT)
        INSERT INTO @VerifiedAttributes
        SELECT Attribute, Rating
        FROM dbo.VerifyAndReturnAttributes(@Attributes)

        -- Debugging: Print contents of @VerifiedAttributes
        PRINT 'Verified Attributes:'
        SELECT * FROM @VerifiedAttributes

        -- Ensure attributes exist in at least one of the attribute tables before inserting
        IF EXISTS (
            SELECT 1
            FROM @VerifiedAttributes AS VA
            WHERE NOT EXISTS (SELECT 1 FROM Goalkeeping_ATT WHERE att_id = VA.Attribute)
              AND NOT EXISTS (SELECT 1 FROM Mental_ATT WHERE att_id = VA.Attribute)
              AND NOT EXISTS (SELECT 1 FROM Physical_ATT WHERE att_id = VA.Attribute)
        )
        BEGIN
            PRINT 'One or more attributes do not exist in the attribute tables.'
            ROLLBACK TRANSACTION;
            THROW 50002, 'One or more attributes do not exist in the attribute tables.', 1;
        END

        -- Insert into GoalkeeperAttributeRating
        INSERT INTO GoalkeeperAttributeRating (att_id, player_id, rating)
        SELECT Attribute, @PlayerID, Rating
        FROM @VerifiedAttributes

        -- Verify all attributes from Goalkeeping_ATT, Mental_ATT, and Physical_ATT are present
        DECLARE @MissingGoalkeepingAtt TABLE (att_id NVARCHAR(255))
        DECLARE @MissingMentalAtt TABLE (att_id NVARCHAR(255))
        DECLARE @MissingPhysicalAtt TABLE (att_id NVARCHAR(255))

        -- Find missing goalkeeping attributes
        INSERT INTO @MissingGoalkeepingAtt
        SELECT att_id
        FROM Goalkeeping_ATT
        WHERE att_id NOT IN (SELECT Attribute FROM @VerifiedAttributes)

        -- Find missing mental attributes
        INSERT INTO @MissingMentalAtt
        SELECT att_id
        FROM Mental_ATT
        WHERE att_id NOT IN (SELECT Attribute FROM @VerifiedAttributes)

        -- Find missing physical attributes
        INSERT INTO @MissingPhysicalAtt
        SELECT att_id
        FROM Physical_ATT
        WHERE att_id NOT IN (SELECT Attribute FROM @VerifiedAttributes)

        -- Print the contents of the missing attribute tables
        PRINT 'Missing Goalkeeping Attributes:'
        SELECT * FROM @MissingGoalkeepingAtt

        PRINT 'Missing Mental Attributes:'
        SELECT * FROM @MissingMentalAtt

        PRINT 'Missing Physical Attributes:'
        SELECT * FROM @MissingPhysicalAtt

        -- Check if any attributes are missing
        IF (SELECT COUNT(*) FROM @MissingGoalkeepingAtt) > 0
           OR (SELECT COUNT(*) FROM @MissingMentalAtt) > 0
           OR (SELECT COUNT(*) FROM @MissingPhysicalAtt) > 0
        BEGIN
            -- Display missing attributes
            SELECT 'Missing Goalkeeping Attributes' AS AttributeType, att_id AS MissingAttribute FROM @MissingGoalkeepingAtt
            UNION ALL
            SELECT 'Missing Mental Attributes' AS AttributeType, att_id AS MissingAttribute FROM @MissingMentalAtt
            UNION ALL
            SELECT 'Missing Physical Attributes' AS AttributeType, att_id AS MissingAttribute FROM @MissingPhysicalAtt

            ROLLBACK TRANSACTION;
            THROW 50003, 'Not all required attributes were inserted.', 1;
        END

        COMMIT TRANSACTION;
    END TRY
    BEGIN CATCH
        IF @@TRANCOUNT > 0
            ROLLBACK TRANSACTION;
        THROW;
    END CATCH
END
GO
