// Copyright 2018 The Ebiten Authors
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
