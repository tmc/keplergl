package keplergl

// Config is the type that describes Kepler.gl configuration.
type Config struct {
	Datasets []*struct {
		Version string `json:"version"`
		Data    struct {
			ID      string          `json:"id"`
			Label   string          `json:"label"`
			Color   []int           `json:"color"`
			AllData [][]interface{} `json:"allData"`
			Fields  []*struct {
				Name         string `json:"name"`
				Type         string `json:"type"`
				Format       string `json:"format"`
				AnalyzerType string `json:"analyzerType"`
			} `json:"fields"`
		} `json:"data"`
	} `json:"datasets"`
	Config *struct {
		Version string `json:"version"`
		Config  *struct {
			VisState struct {
				Filters []interface{} `json:"filters"`
				Layers  []*struct {
					ID     string `json:"id"`
					Type   string `json:"type"`
					Config *struct {
						DataID    string            `json:"dataId"`
						Label     string            `json:"label"`
						Color     []int             `json:"color"`
						Columns   map[string]string `json:"columns"`
						IsVisible bool              `json:"isVisible"`
						VisConfig *struct {
							Opacity          float64 `json:"opacity"`
							StrokeOpacity    float64 `json:"strokeOpacity"`
							Thickness        float64 `json:"thickness"`
							StrokeColor      []int   `json:"strokeColor"`
							ColorAggregation string  `json:"colorAggregation"`
							ColorRange       *struct {
								Name     string   `json:"name"`
								Type     string   `json:"type"`
								Category string   `json:"category"`
								Colors   []string `json:"colors"`
							} `json:"colorRange"`
							StrokeColorRange struct {
								Name     string   `json:"name"`
								Type     string   `json:"type"`
								Category string   `json:"category"`
								Colors   []string `json:"colors"`
							} `json:"strokeColorRange"`
							Radius              int     `json:"radius"`
							SizeRange           []int   `json:"sizeRange"`
							SizeAggregation     string  `json:"sizeAggregation"`
							RadiusRange         []int   `json:"radiusRange"`
							HeightRange         []int   `json:"heightRange"`
							FixedRadius         bool    `json:"fixedRadius"`
							ElevationScale      int     `json:"elevationScale"`
							ElevationPercentile []int   `json:"elevationPercentile"`
							Percentile          []int   `json:"percentile"`
							Coverage            float64 `json:"coverage"`
							Stroked             bool    `json:"stroked"`
							Outline             bool    `json:"outline"`
							Filled              bool    `json:"filled"`
							Enable3D            bool    `json:"enable3d"`
							Wireframe           bool    `json:"wireframe"`
							WorldUnitSize       float64 `json:"worldUnitSize"`
						} `json:"visConfig"`
						Hidden    bool `json:"hidden"`
						TextLabel []struct {
							Field     interface{} `json:"field"`
							Color     []int       `json:"color"`
							Size      int         `json:"size"`
							Offset    []int       `json:"offset"`
							Anchor    string      `json:"anchor"`
							Alignment string      `json:"alignment"`
						} `json:"textLabel"`
					} `json:"config"`
					VisualChannels *struct {
						ColorField       interface{} `json:"colorField"`
						ColorScale       string      `json:"colorScale"`
						StrokeColorField interface{} `json:"strokeColorField"`
						StrokeColorScale string      `json:"strokeColorScale"`
						SizeField        interface{} `json:"sizeField"`
						SizeScale        string      `json:"sizeScale"`
						HeightField      interface{} `json:"heightField"`
						HeightScale      string      `json:"heightScale"`
						RadiusField      interface{} `json:"radiusField"`
						RadiusScale      string      `json:"radiusScale"`
						Thickness        float64     `json:"thickness"`
					} `json:"visualChannels"`
				} `json:"layers"`
				InteractionConfig *struct {
					Tooltip struct {
						FieldsToShow map[string][]struct {
							Name   string      `json:"name"`
							Format interface{} `json:"format"`
						} `json:"fieldsToShow"`
						CompareMode bool   `json:"compareMode"`
						CompareType string `json:"compareType"`
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
				LayerBlending   string        `json:"layerBlending"`
				SplitMaps       []interface{} `json:"splitMaps"`
				AnimationConfig *struct {
					CurrentTime interface{} `json:"currentTime"`
					Speed       int         `json:"speed"`
				} `json:"animationConfig"`
			} `json:"visState"`
			MapState *struct {
				Bearing    float64 `json:"bearing"`
				DragRotate bool    `json:"dragRotate"`
				Latitude   float64 `json:"latitude"`
				Longitude  float64 `json:"longitude"`
				Pitch      float64 `json:"pitch"`
				Zoom       float64 `json:"zoom"`
				IsSplit    bool    `json:"isSplit"`
			} `json:"mapState"`
			MapStyle *struct {
				StyleType      string `json:"styleType"`
				TopLayerGroups struct {
				} `json:"topLayerGroups"`
				VisibleLayerGroups *struct {
					Label          bool `json:"label"`
					Road           bool `json:"road"`
					Border         bool `json:"border"`
					Building       bool `json:"building"`
					Water          bool `json:"water"`
					Land           bool `json:"land"`
					ThreeDBuilding bool `json:"3d building"`
				} `json:"visibleLayerGroups"`
				ThreeDBuildingColor []float64 `json:"threeDBuildingColor"`
				MapStyles           struct {
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
