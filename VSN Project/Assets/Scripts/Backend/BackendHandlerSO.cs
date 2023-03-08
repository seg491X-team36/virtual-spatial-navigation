using System;
using UnityEngine;
using UnityEngine.Networking;
using UnityEngine.Events;
using Newtonsoft.Json;

namespace Backend
{
    public class Handler<Request, Response>
    {
        public void send(Request request, UnityEvent<Response> responses, string address, string token, Action<Response> callback = null)
        {
            var data = JsonConvert.SerializeObject(request);

            var httpRequest = UnityWebRequest.Put(address, data);
            httpRequest.method = "POST"; // work around UnityWebRequest.Post not working
            httpRequest.SetRequestHeader("Content-Type", "application/json");
            httpRequest.SetRequestHeader("token", token);

            Debug.Log(address);

            httpRequest.SendWebRequest().completed += operation =>
            {
                var response = JsonConvert.DeserializeObject<Response>(httpRequest.downloadHandler.text);
                if (callback != null)
                {
                    callback(response);
                }

                responses.Invoke(response);
            };
        }
    }


    [CreateAssetMenu(fileName = "BackendHandlerSO", menuName = "ScriptableObjects/BackendHandlerSO")]
    public class BackendHandlerSO : ScriptableObject
    {
        public BackendEventsSO events;
        public string address; // "http://localhost:3000" or the deployed addresss
        private string token; // the jwt to authenticate the user

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
            var handler = new Handler<SubmitEmailRequest, SubmitEmailResponse> { };
            handler.send(request, events.OnSubmitEmailResponse, address + "/verify-email", "");
        }

        private void handleVerificationRequest(VerificationRequest request)
        {
            var handler = new Handler<VerificationRequest, VerificationResponse> { };
            handler.send(request, events.OnVerificationResponse, address + "/verify-code", "", (response) =>
            {
                // update the token if verification was successful
                if (response.error == null)
                {
                    token = response.token;
                }
            });
        }

        private void handlePendingRequest(PendingRequest request)
        {
            var handler = new Handler<PendingRequest, PendingResponse> { };
            handler.send(request, events.OnPendingResponse, address + "/pending", token);
        }

        private void handleStartExperimentRequest(StartExperimentRequest request)
        {
            var handler = new Handler<StartExperimentRequest, StartExperimentResponse> { };
            handler.send(request, events.OnStartExperimentResponse, address + "/start", token);
        }

        private void handleStartRoundRequest(StartRoundRequest request)
        {
            var handler = new Handler<StartRoundRequest, StartRoundResponse> { };
            handler.send(request, events.OnStartRoundResponse, address + "/round/start", token);
        }

        private void handleStopRoundRequest(StopRoundRequest request)
        {
            var handler = new Handler<StopRoundRequest, StopRoundResponse> { };
            handler.send(request, events.OnStopRoundResponse, address + "/round/stop", token);
        }

        private void handleRecordRequest(RecordRequest request)
        {
            var handler = new Handler<RecordRequest, RecordResponse> { };
            handler.send(request, events.OnRecordResponse, address + "/record", token);
        }
    }
}