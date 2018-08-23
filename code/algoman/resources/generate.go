// Copyright 2018 The Algoman Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate file2byteslice -package=fonts -input=./fonts/arcade_n.ttf -output=./fonts/arcaden.go -var=ArcadeN_ttf
//go:generate file2byteslice -package=fonts -input=./fonts/mplus-1p-regular.ttf -output=./fonts/mplus1pregular.go -var=MPlus1pRegular_ttf
//go:generate file2byteslice -package=fonts -input=./fonts/wqy-zenhei.ttf -output=./fonts/wqyzenhei.go -var=WQYZenHei_ttf

//go:generate file2byteslice -package=images -input=./images/ui.png -output=./images/ui.go -var=UI_png

//go:generate file2byteslice -package=pngs -input=./pngs/fast_forward_512.png -output=./pngs/fast_forward_512.go -var=FAST_FORWARD_512_png
//go:generate file2byteslice -package=pngs -input=./pngs/fast_forward_512_active.png -output=./pngs/fast_forward_512_active.go -var=FAST_FORWARD_512_ACTIVE_png
//go:generate file2byteslice -package=pngs -input=./pngs/pause_512.png -output=./pngs/pause_512.go -var=PAUSE_512_png
//go:generate file2byteslice -package=pngs -input=./pngs/pause_512_active.png -output=./pngs/pause_512_active.go -var=PAUSE_512_ACTIVE_png
//go:generate file2byteslice -package=pngs -input=./pngs/play_512.png -output=./pngs/play_512.go -var=PLAY_512_png
//go:generate file2byteslice -package=pngs -input=./pngs/play_512_active.png -output=./pngs/play_512_active.go -var=PLAY_512_ACTIVE_png
//go:generate file2byteslice -package=pngs -input=./pngs/play_next_512.png -output=./pngs/play_next_512.go -var=PLAY_NEXT_512_png
//go:generate file2byteslice -package=pngs -input=./pngs/play_next_512_active.png -output=./pngs/play_next_512_active.go -var=PLAY_NEXT_512_ACTIVE_png
//go:generate file2byteslice -package=pngs -input=./pngs/previous_track_512.png -output=./pngs/previous_track_512.go -var=PREVIOUS_TRACK_512_png
//go:generate file2byteslice -package=pngs -input=./pngs/previous_track_512_active.png -output=./pngs/previous_track_512_active.go -var=PREVIOUS_TRACK_512_ACTIVE_png
//go:generate file2byteslice -package=pngs -input=./pngs/replay_512.png -output=./pngs/replay_512.go -var=REPLAY_512_png
//go:generate file2byteslice -package=pngs -input=./pngs/replay_512_active.png -output=./pngs/replay_512_active.go -var=REPLAY_512_active_png
//go:generate file2byteslice -package=pngs -input=./pngs/rewind_512.png -output=./pngs/rewind_512.go -var=REWIND_512_png
//go:generate file2byteslice -package=pngs -input=./pngs/rewind_512_active.png -output=./pngs/rewind_512_active.go -var=REWIND_512_active_png

//go:generate ./svg2png -input=./images/backward_off -package=images
//go:generate ./svg2png -input=./images/backward_on -package=images
//go:generate ./svg2png -input=./images/forward_off -package=images
//go:generate ./svg2png -input=./images/forward_on -package=images
//go:generate ./svg2png -input=./images/pause_off -package=images
//go:generate ./svg2png -input=./images/pause_on -package=images
//go:generate ./svg2png -input=./images/play_off -package=images
//go:generate ./svg2png -input=./images/play_on -package=images
//go:generate ./svg2png -input=./images/redo_off -package=images
//go:generate ./svg2png -input=./images/redo_on -package=images
//go:generate ./svg2png -input=./images/step_backward_off -package=images
//go:generate ./svg2png -input=./images/step_backward_on -package=images
//go:generate ./svg2png -input=./images/step_forward_off -package=images
//go:generate ./svg2png -input=./images/step_forward_on -package=images

//go:generate gofmt -s -w .

package resources
