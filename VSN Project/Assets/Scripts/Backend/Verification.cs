using System;

// Types used for verifying users
#nullable enable
namespace Backend
{
    [Serializable]
    public struct SubmitEmailRequest
    {
        public string email;
    }

    [Serializable]
    public struct SubmitEmailResponse
    {
    }

    [Serializable]
    public struct VerificationRequest
    {
        public string email;
        public string code;
    }

    [Serializable]
    public struct VerificationResponse
    {
        public string token;
        public string? error;
    }
}


