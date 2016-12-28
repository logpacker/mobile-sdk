using System;
using System.Net;
using System.Threading;
using System.Text;
using System.IO;
using Newtonsoft.Json;

namespace logpackermobilesdk
{
	public class Client
	{
		public string ClusterURL;
		public string Environment;
		public string Agent;
		public string CloudKey;

		public Client(string clusterURL, string environment, string agent, string cloudKey) {
			if (agent == "") {
				agent = "windowsphone";
			}
			if (environment == "") {
				environment = "development";
			}

			if (clusterURL == "") {
				throw new ArgumentException ("ClusterURL must contain host:port for your LogPacker Cluster");
			}

			// Get response code from LogPacker cluster
			WebRequest request = WebRequest.Create(clusterURL+"/version");
			AutoResetEvent autoResetEvent = new AutoResetEvent(false);
			IAsyncResult asyncResult = request.BeginGetResponse(r => autoResetEvent.Set(), null);
			// Wait until the call is finished
			autoResetEvent.WaitOne();
			HttpWebResponse response = request.EndGetResponse(asyncResult) as HttpWebResponse;
			int code = (int)response.StatusCode;
			if (code != 200) {
				throw new WebException ("ClusterURL "+clusterURL+" isn't reachable");
			}

			ClusterURL = clusterURL;
			Environment = environment;
			Agent = agent;
			CloudKey = cloudKey;
		}

		public Response Send(Event e)
		{
			Validate (e);
			string json = GenerateRequestJSONString (e);

			// Make POST request
			WebRequest request = WebRequest.Create(ClusterURL+"/save");
			request.ContentType = "text/json";
			request.Method = "POST";

			// Set JSON data
			AutoResetEvent autoResetEventStream = new AutoResetEvent(false);
			IAsyncResult asyncResultStream = request.BeginGetRequestStream(r => autoResetEventStream.Set(), null);
			// Wait until tit's finished
			autoResetEventStream.WaitOne();
			Stream stream = request.EndGetRequestStream(asyncResultStream) as Stream;
			StreamWriter streamWriter = new StreamWriter (stream);
			streamWriter.Write(json);
			streamWriter.Flush();

			// Send and wait
			AutoResetEvent autoResetEvent = new AutoResetEvent(false);
			IAsyncResult asyncResult = request.BeginGetResponse(r => autoResetEvent.Set(), null);
			// Wait until the call is finished
			autoResetEvent.WaitOne();
			HttpWebResponse response = request.EndGetResponse(asyncResult) as HttpWebResponse;
			int code = (int)response.StatusCode;
			if (code != 200) {
				throw new WebException (code.ToString() +": "+ClusterURL+"/save isn't reachable");
			}

			// Parse JSON body
			string body;
			using (var sr = new StreamReader(response.GetResponseStream())) {
				body = sr.ReadToEnd();
			}

			// Decode JSON string into Response object
			Response decoded = JsonConvert.DeserializeObject<Response>(body);

			return decoded;
		}

		public Response SendException(Exception e)
		{
			Event ev = new Event(e.Message, e.StackTrace, Event.FatalLogLevel, "", "");
			return Send (ev);
		}

		private bool Validate(Event e)
		{
			if (e.Message == "") {
				throw new ArgumentException("Message cannot be empty");
			}
			if (e.LogLevel < Event.FatalLogLevel || e.LogLevel > Event.NoticeLogLevel) {
				throw new ArgumentException("LogLevel is invalid. Valid are: "+Event.FatalLogLevel.ToString()+" - "+Event.NoticeLogLevel.ToString());
			}

			return true;
		}

		// Request JSON for LogPacker
		private string GenerateRequestJSONString(Event e)
		{
			// We send only one message(Event) per request
			return "{" +
				"\"client\":{" +
					"\"user_id\":\""+e.UserID+"\","+
					"\"user_name\":\""+e.UserName+"\","+
					"\"platform\":\"mobile\","+
					"\"environment\":\""+Environment+"\","+
					"\"agent\":\""+Agent+"\","+
					"\"os\":\"windows\""+
				"},"+
				"\"messages\":[{"+
					"\"message\":\""+e.Message+"\","+
					"\"source\":\""+e.Source+"\","+
					"\"line\":0,"+
					"\"column\":0,"+
					"\"log_level\":"+e.LogLevel.ToString()+","+
					"\"tag_name\":\"windowsphone\""+
				"}],"+
				"\"cloud_key\":\""+CloudKey+"\""+
			"}";
		}
	}

	public class Event
	{
		public string Message;
		public string Source;
		public int LogLevel;
		public string UserID;
		public string UserName;

		// LogLevel constants
		public const int FatalLogLevel = 0;
		public const int ErrorLogLevel = 1;
		public const int WarnLogLevel = 2;
		public const int InfoLogLevel = 3;
		public const int DebugLogLevel = 4;
		public const int NoticeLogLevel = 5;

		public Event(string message, string source, int level, string userID, string userName) {
			Message = message;
			Source = source;
			LogLevel = level;
			UserID = userID;
			UserName = userName;
		}
	}

	// Response format from LogPacker Cluster
	public struct Response
	{
		public string Code;
		public string Error;
		public string[] Messages;
	}
}
