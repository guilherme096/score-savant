USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE dbo.AddGoalkeeperAttributeRating
(
    @PlayerID INT,
    @Attributes NVARCHAR(MAX)
)
AS
BEGIN
    SET NOCOUNT ON;
    BEGIN TRY
        BEGIN TRANSACTION;

        DECLARE @VerifiedAttributes TABLE (Attribute NVARCHAR(255), Rating INT)
        INSERT INTO @VerifiedAttributes
        SELECT Attribute, Rating
        FROM dbo.VerifyAndReturnAttributes(@Attributes)

        -- Insert into GoalkeeperAttributeRating
        INSERT INTO GoalkeeperAttributeRating (att_id, player_id, rating)
        SELECT Attribute, @PlayerID, Rating
        FROM @VerifiedAttributes

        -- Verify insertion for each category
        DECLARE @GoalkeepingCount INT, @MentalCount INT, @PhysicalCount INT

        SELECT @GoalkeepingCount = COUNT(*) FROM Goalkeeping_ATT WHERE att_id IN (SELECT Attribute FROM @VerifiedAttributes)
        SELECT @MentalCount = COUNT(*) FROM Mental_ATT WHERE att_id IN (SELECT Attribute FROM @VerifiedAttributes)
        SELECT @PhysicalCount = COUNT(*) FROM Physical_ATT WHERE att_id IN (SELECT Attribute FROM @VerifiedAttributes)

        IF @GoalkeepingCount = (SELECT COUNT(*) FROM Goalkeeping_ATT) AND
           @MentalCount = (SELECT COUNT(*) FROM Mental_ATT) AND
           @PhysicalCount = (SELECT COUNT(*) FROM Physical_ATT)
        BEGIN
            COMMIT TRANSACTION;
        END
        ELSE
        BEGIN
            ROLLBACK TRANSACTION;
            THROW 50001, 'Not all required attributes were inserted.', 1;
        END
    END TRY
    BEGIN CATCH
        ROLLBACK TRANSACTION;
        THROW;
    END CATCH
END
