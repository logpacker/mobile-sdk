using System;
using System.Net;
using System.Threading;

namespace logpackermobilesdk
{
	public struct Client
	{
		public string ClusterURL;
		public string Environment;
		public string Agent;

		public Client(string clusterURL, string environment, string agent) {
			ClusterURL = clusterURL;
			Environment = environment;
			Agent = agent;
		}
	}

	public class LogPackerMobileSDK
	{
		public LogPackerMobileSDK ()
		{
		}

		public Client NewClient(string clusterURL, string environment, string agent)
		{
			if (agent == "") {
				agent = "mobile";
			}
			if (environment == "") {
				environment = "development";
			}

			if (clusterURL == "") {
				throw new ArgumentException ("ClusterURL must contain host:port for your LogPacker Cluster");
			}

			// Get response code from LogPacker cluster
			AutoResetEvent autoResetEvent = new AutoResetEvent(false);
			WebRequest request = WebRequest.Create(clusterURL+"/version");
			IAsyncResult asyncResult = request.BeginGetResponse(r => autoResetEvent.Set(), null);
			// Wait until the call is finished
			autoResetEvent.WaitOne();
			HttpWebResponse response = request.EndGetResponse(asyncResult) as HttpWebResponse;
			int code = (int)response.StatusCode;
			if (code != 200) {
				throw new WebException ("ClusterURL isn't reachable");
			}

			Client client = new Client (clusterURL, environment, agent);

			return client;
		}
	}
}

