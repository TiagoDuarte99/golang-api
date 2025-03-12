INSERT INTO teams (id, name, country, coach_id, pts, created_at) VALUES
(1, 'Team Alpha', 'Portugal', 8, 0, '2025-02-23 10:56:39'),
(2, 'Team Beta', 'Brazil', 9, 0, '2025-02-23 18:51:04'),
(3, 'Team Gamma', 'Germany', 4, 0, '2025-02-24 22:14:03'),
(4, 'Team Delta', 'Spain', 10, 0, '2025-02-24 22:18:16');


INSERT INTO players (id, name, position, team_id, age, height, weight, goals, assists, created_at) VALUES
(1, 'Player 1', 'Forward', 1, 25, 1.80, 75, 2, 1, '2025-02-23 10:56:39'),
(2, 'Player 2', 'Midfielder', 2, 24, 1.75, 70, 1, 2, '2025-02-23 18:51:04'),
(3, 'Player 3', 'Defender', 1, 26, 1.85, 80, 0, 1, '2025-02-24 22:14:03'),
(4, 'Player 4', 'Goalkeeper', 2, 30, 1.90, 85, 0, 0, '2025-02-24 22:18:16');


INSERT INTO matches (id, home_team_id, away_team_id, date, location, home_goals, away_goals, created_at) 
VALUES (1, 1, 2, '2025-02-25 20:00:00', 'Estádio Central', 3, 2, '2025-02-23 10:56:39');


INSERT INTO match_scorers (match_id, player_id, scored_goals) VALUES
(1, 1, 2), -- Player 1 marcou 2 gols
(1, 2, 1); -- Player 2 marcou 1 gol

INSERT INTO match_assistants (match_id, player_id, assisted_goals) VALUES
(1, 3, 1), -- Player 3 deu 1 assistência
(1, 4, 1); -- Player 4 deu 1 assistência