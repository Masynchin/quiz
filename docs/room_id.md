# Room ID

This document tells what is Room ID: what its purpose, constaints,
how Room ID generates.

## Purpose

Room ID is the unique ID of each active room.

> Active rooms - rooms that created and not played its game yet.

The purpose of Room ID is to uniquely identify rooms. It allows players
to join rooms by their ID.

## Constaints

Every room of all active rooms has the unique ID. This means that when game
of room is played its ID allows to be used for new rooms.

## Generatation

Room ID generates this way:

- First 5 digits is Owner's UserID. If Owner's UserID less than 10_000,
  then leading zeros prepends to fill up to 5 digits.
- Last 2 digits generates randomly.
