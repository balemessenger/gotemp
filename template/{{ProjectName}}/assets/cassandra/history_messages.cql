CREATE TABLE IF NOT EXISTS history_messages (
    user_id int,
    ex_peer_type int,
    peer_id int,
    peer_type int,
    partition_nr int,
    sequence_nr int,
    timebucket int,
    is_deleted boolean,
    PRIMARY KEY ((user_id, peer_id, ex_peer_type, partition_nr), sequence_nr, timebucket)
) WITH CLUSTERING ORDER BY (sequence_nr DESC, timebucket DESC);