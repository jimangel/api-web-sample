variable "web_image" {}

resource "kubernetes_manifest" "cloudrun-web" {

  manifest = {
    "apiVersion" : "serving.knative.dev/v1",
    "kind" : "Service",
    "metadata" : {
      "name" : "web-app-${var.environment}",
      "namespace" : var.environment
    },
    "spec" : {
      "template" : {
        "spec" : {
          "timeoutSeconds" : 300,
          "containers" : [
            {
              "name" : "user-container",
              "image" : var.web_image,
              "ports" : [
                {
                  "containerPort" : 8080,
                  "protocol" : "TCP"
                }
              ],
              "env" : [
                {
                  "name" : "WEB_PORT",
                  "value" : "8080"
                },
                {
                  "name" : "API_PORT",
                  "value" : "80"
                },
                {
                  "name" : "API_URL",
                  "value" : "api-app-${var.environment}.${var.environment}.svc.cluster.local"
                },
                {
                  "name" : "API_HTTP_S",
                  "value" : "http"
                },
                {
                  "name" : "REQUEST_HOST",
                  "value" : "api-app-${var.environment}.default.svc.cluster.local"
                }
              ],
              "resources" : {
                "limits" : {
                  "memory" : "256Mi"
                }
              },
              "readinessProbe" : {
                "successThreshold" : 1,
                "tcpSocket" : {}
              }
            }
          ]
        }
      },
      "traffic" : [
        {
          "percent" : 100,
          "latestRevision" : true
        }
      ]
    }
  }
}
