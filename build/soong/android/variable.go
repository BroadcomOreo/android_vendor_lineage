package android
type Product_variables struct {	
        Disable_ashmem_tracking struct {
		Cflags []string
        }
	Egl_workaround_bug_10194508 struct {
		Cppflags []string
	}
	Has_legacy_camera_hal1 struct {
		Cflags []string
	}
	Needs_legacy_camera_hal1_dyn_native_handle struct {
		Cppflags []string
	}
	Needs_text_relocations struct {
		Cppflags []string
	}
	Target_shim_libs struct {
		Cppflags []string
	}
	Uses_generic_camera_parameter_library struct {
		Srcs []string
	}
	Uses_nvidia_enhancements struct {
		Cppflags []string
	}
	Uses_qcom_bsp_legacy struct {
		Cppflags []string
	}
        Force_screenshot_cpu_path struct {
                Cppflags []string
        }	
	Uses_qti_camera_device struct {
		Cppflags []string
		Shared_libs []string
	}
}

type ProductVariables struct {
	Force_screenshot_cpu_path    *bool `json:",omitempty"`	
	Disable_ashmem_tracking    *bool `json:",omitempty"`
	Egl_workaround_bug_10194508    *bool `json:",omitempty"`
	Has_legacy_camera_hal1  *bool `json:",omitempty"`
	Needs_legacy_camera_hal1_dyn_native_handle  *bool `json:",omitempty"`
	Needs_text_relocations  *bool `json:",omitempty"`
	Specific_camera_parameter_library  *string `json:",omitempty"`
	Target_shim_libs  *string `json:",omitempty"`
	Uses_generic_camera_parameter_library  *bool `json:",omitempty"`
	Uses_nvidia_enhancements  *bool `json:",omitempty"`
	Uses_qcom_bsp_legacy  *bool `json:",omitempty"`
	Uses_qti_camera_device  *bool `json:",omitempty"`
}
