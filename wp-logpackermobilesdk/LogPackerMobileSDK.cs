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

	public struct Event
	{
		public string Message;
		public string Source;
		public int LogLevel;
		public string UserID;
		public string UserName;

		public Event(string message, string source, int level, string userID, string userName) {
			Message = message;
			Source = source;
			LogLevel = level;
			UserID = userID;
			UserName = userName;
		}
	}

	public class LogPackerMobileSDK
	{
		// LogLevel constants
		public const int FatalLogLevel = 0;
		public const int ErrorLogLevel = 1;
		public const int WarnLogLevel = 2;
		public const int InfoLogLevel = 3;
		public const int DebugLogLevel = 4;
		public const int NoticeLogLevel = 5;

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

		public bool Send(Event e)
		{
			Validate (e);
			string json = GenerateRequestJSONString (e);
			return true;
		}

		private bool Validate(Event e)
		{
			if (e.Message == "") {
				throw new ArgumentException("Message cannot be empty");
			}
			if (e.LogLevel < FatalLogLevel || e.LogLevel > NoticeLogLevel) {
				throw new ArgumentException("LogLevel is invalid. Valid are: "+FatalLogLevel+" - "+NoticeLogLevel);
			}

			return true;
		}

		private string GenerateRequestJSONString(Event e)
		{
			return @"{
				""client"":{},
				""messages"":[]
			}";
		}
	}
}

