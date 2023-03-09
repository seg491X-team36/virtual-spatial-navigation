using UnityEngine;
using Backend;

namespace Backend
{
    [CreateAssetMenu(fileName = "DummyHandlerSO", menuName = "ScriptableObjects/DummyHandlerSO")]
    public class DummyHandlerSO : ScriptableObject
    {
        public BackendEventsSO events;

        private void OnEnable()
        {
            // add listeners to "request" events
            events.OnSubmitEmailRequest.AddListener(handleSubmitEmailRequest);
            events.OnVerificationRequest.AddListener(handleVerificationRequest);
            events.OnPendingRequest.AddListener(handlePendingRequest);
            events.OnStartExperimentRequest.AddListener(handleStartExperimentRequest);
            events.OnStartRoundRequest.AddListener(handleStartRoundRequest);
            events.OnStopRoundRequest.AddListener(handleStopRoundRequest);
            events.OnRecordRequest.AddListener(handleRecordRequest);
        }

        private void OnDisable()
        {
            // remove listeners from "request" events
            events.OnSubmitEmailRequest.RemoveListener(handleSubmitEmailRequest);
            events.OnVerificationRequest.RemoveListener(handleVerificationRequest);
            events.OnPendingRequest.RemoveListener(handlePendingRequest);
            events.OnStartExperimentRequest.RemoveListener(handleStartExperimentRequest);
            events.OnStartRoundRequest.RemoveListener(handleStartRoundRequest);
            events.OnStopRoundRequest.RemoveListener(handleStopRoundRequest);
            events.OnRecordRequest.RemoveListener(handleRecordRequest);
        }

        private void handleSubmitEmailRequest(SubmitEmailRequest request)
        {
        }

        private void handleVerificationRequest(VerificationRequest request)
        {

        }

        private void handlePendingRequest(PendingRequest request)
        {
        }

        private void handleStartExperimentRequest(StartExperimentRequest request)
        {
        }

        private void handleStartRoundRequest(StartRoundRequest request)
        {
        }

        private void handleStopRoundRequest(StopRoundRequest request)
        {
        }

        private void handleRecordRequest(RecordRequest request)
        {
        }
    }
}