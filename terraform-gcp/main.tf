resource "random_id" "project_suffix" {
  byte_length = 2
}

resource "time_sleep" "wait_for_billing" {
  depends_on      = [google_billing_project_info.project-billing]
  create_duration = "120s"
}

# Creating custom project for red team testing
resource "google_project" "model-armor-c0d3x" {
  name            = "model-armor-c0d3x"
  project_id      = "model-armor-c0d3x-${random_id.project_suffix.hex}"
  deletion_policy = "DELETE"
  lifecycle {
    ignore_changes = [billing_account]
  }
}

# Attaching Billing resource to model armor C0D3X project 
resource "google_billing_project_info" "project-billing" {
  project         = google_project.model-armor-c0d3x.project_id
  billing_account = "01A883-C7E889-8284DA"
}

# Enable API for model armor 
resource "google_project_service" "modelarmor_api" {
  project            = google_project.model-armor-c0d3x.project_id
  service            = "modelarmor.googleapis.com"
  disable_on_destroy = false
}

# Enable API for Model garden aka vertex api
resource "google_project_service" "agent-platform-api" {
  project            = google_project.model-armor-c0d3x.project_id
  service            = "aiplatform.googleapis.com"
  disable_on_destroy = false
}

# Enabling Model Armor within the project
resource "google_model_armor_floorsetting" "model-armor-basic-floorsetting" {
  parent                           = "projects/${google_project.model-armor-c0d3x.project_id}"
  location                         = "global"
  integrated_services              = ["AI_PLATFORM"]
  enable_floor_setting_enforcement = true
  depends_on = [google_project_service.modelarmor_api,
    google_billing_project_info.project-billing,
  time_sleep.wait_for_billing]
  filter_config {
    malicious_uri_filter_settings {
      filter_enforcement = "DISABLED"
    }
    pi_and_jailbreak_filter_settings {
      filter_enforcement = "DISABLED"
    }
    sdp_settings {
      basic_config {
        filter_enforcement = "DISABLED"
      }
    }
    rai_settings {
      rai_filters {
        filter_type = "SEXUALLY_EXPLICIT"
      }
      rai_filters {
        filter_type = "HATE_SPEECH"
      }
      rai_filters {
        filter_type = "HARASSMENT"
      }
      rai_filters {
        filter_type = "DANGEROUS"
      }
    }
  }
  ai_platform_floor_setting {
    inspect_and_block    = true
    enable_cloud_logging = true
  }
}

# Creating model armor templates for the project 
resource "google_model_armor_template" "confidence-low" {
  location    = "europe-west4"
  template_id = "Confidence-level-low"
  project     = google_project.model-armor-c0d3x.project_id
  labels = {
    "model-armor-redteam" = "low-confidence-template"
  }

  filter_config {
    malicious_uri_filter_settings {
      filter_enforcement = "ENABLED"
    }
    pi_and_jailbreak_filter_settings {
      filter_enforcement = "ENABLED"
      confidence_level   = "LOW_AND_ABOVE"
    }
    rai_settings {
      rai_filters {
        filter_type      = "DANGEROUS"
        confidence_level = "LOW_AND_ABOVE"
      }
      rai_filters {
        filter_type      = "SEXUALLY_EXPLICIT"
        confidence_level = "LOW_AND_ABOVE"
      }
      rai_filters {
        filter_type      = "HATE_SPEECH"
        confidence_level = "LOW_AND_ABOVE"
      }
      rai_filters {
        filter_type      = "HARASSMENT"
        confidence_level = "LOW_AND_ABOVE"
      }
    }
    sdp_settings {
      basic_config {
        filter_enforcement = "ENABLED"
      }
    }
  }
  template_metadata {
    log_template_operations = true # log template crud operations.
    log_sanitize_operations = true #  log sanitize operations.
    multi_language_detection {
      enable_multi_language_detection = false
    }
    enforcement_type = "INSPECT_AND_BLOCK"
  }
}
