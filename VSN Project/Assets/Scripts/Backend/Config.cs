using System;

// Types used for the experiment config
#nullable enable
namespace Backend
{
    [Serializable]
    public struct ExperimentConfig
    {
        public int roundsTotal;
        public string resume;
        public int[] spawnSequence;
        public int rewardPosition;
        public Arena arena;
        public ArenaObject[] arenaObjects;
    }

    [Serializable]
    public struct Arena
    {
        public string[] objectNames;
        public Position[] rewardPositions;
        public Position[] spawnPositions;
    }

    [Serializable]
    public struct ArenaObject
    {
        public string objectName;
        public Position position;
    }

    [Serializable]
    public struct Position
    {
        public float x;
        public float y;
        public float z;
    }
}