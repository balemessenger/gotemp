CREATE TABLE IF NOT EXISTS example (
    user_id int,
    sequence_nr int,
    timebucket int,
    is_deleted boolean,
    PRIMARY KEY ((user_id, partition_nr), sequence_nr, timebucket)
) WITH CLUSTERING ORDER BY (sequence_nr DESC, timebucket DESC);