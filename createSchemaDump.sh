#!/bin/bash
pg_dump -U postgres -h localhost -p 5432 -s -F p -v -f reminder_dump.sql reminder
