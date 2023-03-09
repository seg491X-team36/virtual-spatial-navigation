using UnityEngine;
using UnityEngine.Events;

namespace Backend
{
    [CreateAssetMenu(fileName = "BackendEventsSO", menuName = "ScriptableObjects/BackendEventsSO")]
    public class BackendEventsSO : ScriptableObject
    {
        // verification step 1: the user submits their email address to receive a verification code
        public UnityEvent<SubmitEmailRequest> OnSubmitEmailRequest;
        public UnityEvent<SubmitEmailResponse> OnSubmitEmailResponse;

        // verification step 2: the user submits their verification code
        public UnityEvent<VerificationRequest> OnVerificationRequest;

        // verification step 3: response is received
        public UnityEvent<VerificationResponse> OnVerificationResponse;

        // get pending experiment
        public UnityEvent<PendingRequest> OnPendingRequest;
        public UnityEvent<PendingResponse> OnPendingResponse;

        // starting an experiment
        public UnityEvent<StartExperimentRequest> OnStartExperimentRequest;
        public UnityEvent<StartExperimentResponse> OnStartExperimentResponse;

        // starting a round
        public UnityEvent<StartRoundRequest> OnStartRoundRequest;
        public UnityEvent<StartRoundResponse> OnStartRoundResponse;

        // stopping a round
        public UnityEvent<StopRoundRequest> OnStopRoundRequest;
        public UnityEvent<StopRoundResponse> OnStopRoundResponse;

        // recording experiment data
        public UnityEvent<RecordRequest> OnRecordRequest;
        public UnityEvent<RecordResponse> OnRecordResponse;
    }
}