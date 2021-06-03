package keplergl

// Config is the type that describes Kepler.gl configuration.
type Config struct {
	Datasets []struct {
		Version string `json:"version"`
		Data    struct {
			ID      string      `json:"id"`
			Label   string      `json:"label"`
			Color   []int       `json:"color"`
			Alldata [][]float64 `json:"allData"`
			Fields  []struct {
				Name         string `json:"name"`
				Type         string `json:"type"`
				Format       string `json:"format"`
				Analyzertype string `json:"analyzerType"`
			} `json:"fields"`
		} `json:"data"`
	} `json:"datasets"`
	Config struct {
		Version string `json:"version"`
		Config  struct {
			Visstate struct {
				Filters []interface{} `json:"filters"`
				Layers  []struct {
					ID     string `json:"id"`
					Type   string `json:"type"`
					Config struct {
						Dataid  string `json:"dataId"`
						Label   string `json:"label"`
						Color   []int  `json:"color"`
						Columns struct {
							Lat string `json:"lat"`
							Lng string `json:"lng"`
						} `json:"columns"`
						Isvisible bool `json:"isVisible"`
						Visconfig struct {
							Opacity       float64 `json:"opacity"`
							Worldunitsize int     `json:"worldUnitSize"`
							Resolution    int     `json:"resolution"`
							Colorrange    struct {
								Name     string   `json:"name"`
								Type     string   `json:"type"`
								Category string   `json:"category"`
								Colors   []string `json:"colors"`
							} `json:"colorRange"`
							Coverage            int    `json:"coverage"`
							Sizerange           []int  `json:"sizeRange"`
							Percentile          []int  `json:"percentile"`
							Elevationpercentile []int  `json:"elevationPercentile"`
							Elevationscale      int    `json:"elevationScale"`
							Coloraggregation    string `json:"colorAggregation"`
							Sizeaggregation     string `json:"sizeAggregation"`
							Enable3D            bool   `json:"enable3d"`
						} `json:"visConfig"`
						Hidden    bool `json:"hidden"`
						Textlabel []struct {
							Field     interface{} `json:"field"`
							Color     []int       `json:"color"`
							Size      int         `json:"size"`
							Offset    []int       `json:"offset"`
							Anchor    string      `json:"anchor"`
							Alignment string      `json:"alignment"`
						} `json:"textLabel"`
					} `json:"config"`
					Visualchannels struct {
						Colorfield struct {
							Name string `json:"name"`
							Type string `json:"type"`
						} `json:"colorField"`
						Colorscale string      `json:"colorScale"`
						Sizefield  interface{} `json:"sizeField"`
						Sizescale  string      `json:"sizeScale"`
					} `json:"visualChannels"`
				} `json:"layers"`
				Interactionconfig struct {
					Tooltip struct {
						Fieldstoshow struct {
							Bhclc0Usg []struct {
								Name   string      `json:"name"`
								Format interface{} `json:"format"`
							} `json:"bhclc0usg"`
						} `json:"fieldsToShow"`
						Comparemode bool   `json:"compareMode"`
						Comparetype string `json:"compareType"`
						Enabled     bool   `json:"enabled"`
					} `json:"tooltip"`
					Brush struct {
						Size    float64 `json:"size"`
						Enabled bool    `json:"enabled"`
					} `json:"brush"`
					Geocoder struct {
						Enabled bool `json:"enabled"`
					} `json:"geocoder"`
					Coordinate struct {
						Enabled bool `json:"enabled"`
					} `json:"coordinate"`
				} `json:"interactionConfig"`
				Layerblending   string        `json:"layerBlending"`
				Splitmaps       []interface{} `json:"splitMaps"`
				Animationconfig struct {
					Currenttime interface{} `json:"currentTime"`
					Speed       int         `json:"speed"`
				} `json:"animationConfig"`
			} `json:"visState"`
			Mapstate struct {
				Bearing    int     `json:"bearing"`
				Dragrotate bool    `json:"dragRotate"`
				Latitude   float64 `json:"latitude"`
				Longitude  float64 `json:"longitude"`
				Pitch      int     `json:"pitch"`
				Zoom       int     `json:"zoom"`
				Issplit    bool    `json:"isSplit"`
			} `json:"mapState"`
			Mapstyle struct {
				Styletype      string `json:"styleType"`
				Toplayergroups struct {
				} `json:"topLayerGroups"`
				Visiblelayergroups struct {
					Label          bool `json:"label"`
					Road           bool `json:"road"`
					Border         bool `json:"border"`
					Building       bool `json:"building"`
					Water          bool `json:"water"`
					Land           bool `json:"land"`
					ThreeDBuilding bool `json:"3d building"`
				} `json:"visibleLayerGroups"`
				Threedbuildingcolor []float64 `json:"threeDBuildingColor"`
				Mapstyles           struct {
				} `json:"mapStyles"`
			} `json:"mapStyle"`
		} `json:"config"`
	} `json:"config"`
	Info struct {
		App         string `json:"app"`
		CreatedAt   string `json:"created_at"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"info"`
}
