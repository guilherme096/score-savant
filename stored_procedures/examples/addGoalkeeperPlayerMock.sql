USE [p5g5]
GO

-- Example call to AddPlayer stored procedure
EXEC dbo.AddPlayer
    @name = 'Johnny Boy',
    @age = 18,
    @weight = 75,
    @height = 180,
    @nation = 'USA',
    @league = 'MLS',
    @club = 'LA Galaxy',
    @foot = 'Right',
    @value = 100000,
    @player_type = 1, -- 0 for outfield player, 1 for goalkeeper
    @position = 'GK',
    @role = 'Goalkeeper (De)',
    @wage = 5000,
    @contract_end = '2029-06-30',
    @release_clause = 50000,
    @attributes = 'Aerial_Reach:17,Command_Of_Area:15,Communication:16,Eccentricity:8,Gk_First_Touch:12,Gk_Free_Kick_Taking:9,Gk_Passing:13,Gk_Penalty_Taking:10,Gk_Technique:16,Handling:15,Kicking:13,One_On_Ones:18,Punching:5,Reflexes:18,Rushing_Out:12,Throwing:20,Aggression:10,Anticipation:18,Bravery:17,Composure:18,Concentration:15,Decisons:16,Determination:20,Flair:13,Leadership:15,Off_The_Ball:19,Positioning:3,Teamwork:17,Vision:15,Work_Rate:20,Acceleration:17,Agility:14,Balance:14,Jumping_Reach:16,Natural_Fitness:20,Pace:18,Stamina:20,Strength:17'