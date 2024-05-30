USE [p5g5]
GO

CREATE VIEW RandomPlayer AS
SELECT TOP 1 player_id, name, age, weight, height, nation_id, club_id, foot, value, url
FROM Player
ORDER BY NEWID();
GO
