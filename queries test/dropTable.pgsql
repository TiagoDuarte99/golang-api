INSERT INTO matches (home_team_id, away_team_id, date, location, home_goals, away_goals, created_at)
VALUES (2, 3, '2025-03-06T15:00:00Z', 'Stadium A', 2, 1, CURRENT_TIMESTAMP)
RETURNING id;

INSERT INTO match_scorers (match_id, player_id, scored_goals)
VALUES
  (1, 5, 2), -- Jogador 5 marcou 2 gols no jogo 1
  (1, 7, 1);

  INSERT INTO match_assistants (match_id, player_id, assisted_goals)
VALUES
  (1, 8, 1);