using UnityEngine;
using Grpc.Core;
using Schema;

[CreateAssetMenu(menuName = "ScriptableObjects/ExampleScriptableObject")]
public class ServerClientSO : ScriptableObject
{
    private void OnEnable()
    {
        background();
    }

    private async void background()
    {
        var channel = new Channel("localhost:3000", ChannelCredentials.Insecure);
        var req = new CounterRequest { Increment = 1 };
        var client = new Backend.BackendClient(channel);

        var call = client.Counter();
        for (var i = 0; i < 100; i++)
        {
            await call.RequestStream.WriteAsync(req);
        }
        await call.RequestStream.CompleteAsync();
        var res = await call;
        Debug.Log("COUNTER VALUE " + res.Value);
    }
}