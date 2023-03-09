using System;

// Types used for recording experiment data
#nullable enable
namespace Backend
{
    [Serializable]
    public class ExperimentData
    {
        public Frame[] frames = { };
        public Event[] events = { };
    }

    [Serializable]
    public struct RecordRequest
    {
        public ExperimentData data;
    }

    [Serializable]
    public struct RecordResponse
    {
        public string? error;
    }

    [Serializable]
    public struct Frame
    {
        public string timestamp;
        public float x;
        public float y;
        public float z;
        public float xRot;
        public float yRot;
        public float zRot;
    }

    [Serializable]
    public struct Event
    {
        public string name;
        public string timestamp;
    }
}