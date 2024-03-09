-- logs_dummy_data.sql

-- Insert dummy data
INSERT INTO logs (log_level, message, resource_id, trace_id, span_id, commit, metadata, created_at)
VALUES
    ('INFO', 'This is an informational message', 'resource-1', 'trace-1', 'span-1', 'abc123', '{"key1": "value1", "key2": 42}', NOW()),
    ('WARNING', 'This is a warning message', 'resource-2', 'trace-2', 'span-2', 'def456', '{"key3": "value3", "key4": false}', NOW()),
    ('ERROR', 'This is an error message', 'resource-3', 'trace-3', 'span-3', 'ghi789', '{"key5": 3.14, "key6": null}', NOW()),
    ('DEBUG', 'This is a debug message', 'resource-4', 'trace-4', 'span-4', 'jkl012', '{"key7": [1, 2, 3], "key8": {"nested": "value"}}', NOW()),
    ('CRITICAL', 'This is a critical message', 'resource-5', 'trace-5', 'span-5', 'mno345', '{"key9": "value9", "key10": true}', NOW());