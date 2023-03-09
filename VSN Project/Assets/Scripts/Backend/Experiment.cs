using System;

// Types used for the experiment lifecycle
#nullable enable
namespace Backend
{
    [Serializable]
    public struct PendingRequest
    {
        // empty
    }

    [Serializable]
    public struct PendingResponse
    {
        public string? experimentId;
        public bool experimentInProgress;
        public int pending;
    }

    [Serializable]
    public struct StartExperimentRequest
    {
        public string experimentId;
    }

    public struct StartExperimentData
    {
        public ExperimentConfig config;
        public Status status;
        public Frame? frame;
    }

    [Serializable]
    public struct StartExperimentResponse
    {
        public StartExperimentData? data;
        public string? error;
    }

    [Serializable]
    public struct StartRoundRequest
    {
        // empty
    }

    [Serializable]
    public struct StartRoundResponse
    {
        public Status? status;
        public string? error;
    }

    [Serializable]
    public struct StopRoundRequest
    {
        public ExperimentData data;
    }

    [Serializable]
    public struct StopRoundResponse
    {
        public Status? status;
        public string? error;
    }

    [Serializable]
    public struct Status
    {
        public bool roundInProgress;
        public int roundsCompleted;
        public int roundsTotal;
    }
}