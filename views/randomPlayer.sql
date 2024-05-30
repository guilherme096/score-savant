USE [p5g5]
GO

CREATE VIEW RandomPlayer AS
SELECT TOP 1 
    p.player_id, 
    p.name AS player_name, 
    p.nation_id, 
    n.name AS nation_name, 
    p.club_id, 
    c.name AS club_name,
    p.url AS player_url
FROM 
    Player p
JOIN 
    Nation n ON p.nation_id = n.nation_id
JOIN 
    Club c ON p.club_id = c.club_id
ORDER BY 
    NEWID();
GO
