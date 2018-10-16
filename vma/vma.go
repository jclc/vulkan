// The autogenerated code is released under Public Domain.

// WARNING: This file has automatically been generated on Tue, 16 Oct 2018 02:42:19 EEST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package vma

/*
#cgo CFLAGS: -I.. -DVK_NO_PROTOTYPES
#include "../vulkan/vulkan.h"
#include "vk_mem_alloc.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// CreateAllocator function as declared in vma/vk_mem_alloc.h:1701
func CreateAllocator(pCreateInfo *AllocatorCreateInfo, pAllocator *Allocator) VkResult {
	cpCreateInfo, _ := pCreateInfo.PassRef()
	cpAllocator, _ := (*C.VmaAllocator)(unsafe.Pointer(pAllocator)), cgoAllocsUnknown
	__ret := C.vmaCreateAllocator(cpCreateInfo, cpAllocator)
	__v := (VkResult)(__ret)
	return __v
}

// DestroyAllocator function as declared in vma/vk_mem_alloc.h:1706
func DestroyAllocator(allocator Allocator) {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	C.vmaDestroyAllocator(callocator)
}

// GetPhysicalDeviceProperties function as declared in vma/vk_mem_alloc.h:1713
func GetPhysicalDeviceProperties(allocator Allocator, ppPhysicalDeviceProperties **VkPhysicalDeviceProperties) {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cppPhysicalDeviceProperties, _ := ppPhysicalDeviceProperties.PassRef()
	C.vmaGetPhysicalDeviceProperties(callocator, cppPhysicalDeviceProperties)
}

// GetMemoryProperties function as declared in vma/vk_mem_alloc.h:1721
func GetMemoryProperties(allocator Allocator, ppPhysicalDeviceMemoryProperties **VkPhysicalDeviceMemoryProperties) {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cppPhysicalDeviceMemoryProperties, _ := ppPhysicalDeviceMemoryProperties.PassRef()
	C.vmaGetMemoryProperties(callocator, cppPhysicalDeviceMemoryProperties)
}

// GetMemoryTypeProperties function as declared in vma/vk_mem_alloc.h:1731
func GetMemoryTypeProperties(allocator Allocator, memoryTypeIndex uint32, pFlags *uint32) {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cmemoryTypeIndex, _ := (C.uint32_t)(memoryTypeIndex), cgoAllocsUnknown
	cpFlags, _ := (*C.VkMemoryPropertyFlags)(unsafe.Pointer(pFlags)), cgoAllocsUnknown
	C.vmaGetMemoryTypeProperties(callocator, cmemoryTypeIndex, cpFlags)
}

// SetCurrentFrameIndex function as declared in vma/vk_mem_alloc.h:1744
func SetCurrentFrameIndex(allocator Allocator, frameIndex uint32) {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cframeIndex, _ := (C.uint32_t)(frameIndex), cgoAllocsUnknown
	C.vmaSetCurrentFrameIndex(callocator, cframeIndex)
}

// CalculateStats function as declared in vma/vk_mem_alloc.h:1775
func CalculateStats(allocator Allocator, pStats *Stats) {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cpStats, _ := pStats.PassRef()
	C.vmaCalculateStats(callocator, cpStats)
}

// BuildStatsString function as declared in vma/vk_mem_alloc.h:1786
func BuildStatsString(allocator Allocator, ppStatsString []*byte, detailedMap uint32) {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cppStatsString, _ := (**C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&ppStatsString)).Data)), cgoAllocsUnknown
	cdetailedMap, _ := (C.VkBool32)(detailedMap), cgoAllocsUnknown
	C.vmaBuildStatsString(callocator, cppStatsString, cdetailedMap)
}

// FreeStatsString function as declared in vma/vk_mem_alloc.h:1791
func FreeStatsString(allocator Allocator, pStatsString *byte) {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cpStatsString, _ := (*C.char)(unsafe.Pointer(pStatsString)), cgoAllocsUnknown
	C.vmaFreeStatsString(callocator, cpStatsString)
}

// FindMemoryTypeIndex function as declared in vma/vk_mem_alloc.h:2022
func FindMemoryTypeIndex(allocator Allocator, memoryTypeBits uint32, pAllocationCreateInfo *AllocationCreateInfo, pMemoryTypeIndex *uint32) VkResult {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cmemoryTypeBits, _ := (C.uint32_t)(memoryTypeBits), cgoAllocsUnknown
	cpAllocationCreateInfo, _ := pAllocationCreateInfo.PassRef()
	cpMemoryTypeIndex, _ := (*C.uint32_t)(unsafe.Pointer(pMemoryTypeIndex)), cgoAllocsUnknown
	__ret := C.vmaFindMemoryTypeIndex(callocator, cmemoryTypeBits, cpAllocationCreateInfo, cpMemoryTypeIndex)
	__v := (VkResult)(__ret)
	return __v
}

// FindMemoryTypeIndexForBufferInfo function as declared in vma/vk_mem_alloc.h:2040
func FindMemoryTypeIndexForBufferInfo(allocator Allocator, pBufferCreateInfo *VkBufferCreateInfo, pAllocationCreateInfo *AllocationCreateInfo, pMemoryTypeIndex *uint32) VkResult {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cpBufferCreateInfo, _ := pBufferCreateInfo.PassRef()
	cpAllocationCreateInfo, _ := pAllocationCreateInfo.PassRef()
	cpMemoryTypeIndex, _ := (*C.uint32_t)(unsafe.Pointer(pMemoryTypeIndex)), cgoAllocsUnknown
	__ret := C.vmaFindMemoryTypeIndexForBufferInfo(callocator, cpBufferCreateInfo, cpAllocationCreateInfo, cpMemoryTypeIndex)
	__v := (VkResult)(__ret)
	return __v
}

// FindMemoryTypeIndexForImageInfo function as declared in vma/vk_mem_alloc.h:2058
func FindMemoryTypeIndexForImageInfo(allocator Allocator, pImageCreateInfo *VkImageCreateInfo, pAllocationCreateInfo *AllocationCreateInfo, pMemoryTypeIndex *uint32) VkResult {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cpImageCreateInfo, _ := pImageCreateInfo.PassRef()
	cpAllocationCreateInfo, _ := pAllocationCreateInfo.PassRef()
	cpMemoryTypeIndex, _ := (*C.uint32_t)(unsafe.Pointer(pMemoryTypeIndex)), cgoAllocsUnknown
	__ret := C.vmaFindMemoryTypeIndexForImageInfo(callocator, cpImageCreateInfo, cpAllocationCreateInfo, cpMemoryTypeIndex)
	__v := (VkResult)(__ret)
	return __v
}

// CreatePool function as declared in vma/vk_mem_alloc.h:2203
func CreatePool(allocator Allocator, pCreateInfo *PoolCreateInfo, pPool *Pool) VkResult {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cpCreateInfo, _ := pCreateInfo.PassRef()
	cpPool, _ := (*C.VmaPool)(unsafe.Pointer(pPool)), cgoAllocsUnknown
	__ret := C.vmaCreatePool(callocator, cpCreateInfo, cpPool)
	__v := (VkResult)(__ret)
	return __v
}

// DestroyPool function as declared in vma/vk_mem_alloc.h:2210
func DestroyPool(allocator Allocator, pool Pool) {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cpool, _ := *(*C.VmaPool)(unsafe.Pointer(&pool)), cgoAllocsUnknown
	C.vmaDestroyPool(callocator, cpool)
}

// GetPoolStats function as declared in vma/vk_mem_alloc.h:2220
func GetPoolStats(allocator Allocator, pool Pool, pPoolStats *PoolStats) {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cpool, _ := *(*C.VmaPool)(unsafe.Pointer(&pool)), cgoAllocsUnknown
	cpPoolStats, _ := pPoolStats.PassRef()
	C.vmaGetPoolStats(callocator, cpool, cpPoolStats)
}

// MakePoolAllocationsLost function as declared in vma/vk_mem_alloc.h:2231
func MakePoolAllocationsLost(allocator Allocator, pool Pool, pLostAllocationCount *uint) {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cpool, _ := *(*C.VmaPool)(unsafe.Pointer(&pool)), cgoAllocsUnknown
	cpLostAllocationCount, _ := (*C.size_t)(unsafe.Pointer(pLostAllocationCount)), cgoAllocsUnknown
	C.vmaMakePoolAllocationsLost(callocator, cpool, cpLostAllocationCount)
}

// CheckPoolCorruption function as declared in vma/vk_mem_alloc.h:2250
func CheckPoolCorruption(allocator Allocator, pool Pool) VkResult {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cpool, _ := *(*C.VmaPool)(unsafe.Pointer(&pool)), cgoAllocsUnknown
	__ret := C.vmaCheckPoolCorruption(callocator, cpool)
	__v := (VkResult)(__ret)
	return __v
}

// AllocateMemory function as declared in vma/vk_mem_alloc.h:2331
func AllocateMemory(allocator Allocator, pVkMemoryRequirements *VkMemoryRequirements, pCreateInfo *AllocationCreateInfo, pAllocation *Allocation, pAllocationInfo *AllocationInfo) VkResult {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cpVkMemoryRequirements, _ := pVkMemoryRequirements.PassRef()
	cpCreateInfo, _ := pCreateInfo.PassRef()
	cpAllocation, _ := (*C.VmaAllocation)(unsafe.Pointer(pAllocation)), cgoAllocsUnknown
	cpAllocationInfo, _ := pAllocationInfo.PassRef()
	__ret := C.vmaAllocateMemory(callocator, cpVkMemoryRequirements, cpCreateInfo, cpAllocation, cpAllocationInfo)
	__v := (VkResult)(__ret)
	return __v
}

// AllocateMemoryForBuffer function as declared in vma/vk_mem_alloc.h:2344
func AllocateMemoryForBuffer(allocator Allocator, buffer VkBuffer, pCreateInfo *AllocationCreateInfo, pAllocation *Allocation, pAllocationInfo *AllocationInfo) VkResult {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cbuffer, _ := *(*C.VkBuffer)(unsafe.Pointer(&buffer)), cgoAllocsUnknown
	cpCreateInfo, _ := pCreateInfo.PassRef()
	cpAllocation, _ := (*C.VmaAllocation)(unsafe.Pointer(pAllocation)), cgoAllocsUnknown
	cpAllocationInfo, _ := pAllocationInfo.PassRef()
	__ret := C.vmaAllocateMemoryForBuffer(callocator, cbuffer, cpCreateInfo, cpAllocation, cpAllocationInfo)
	__v := (VkResult)(__ret)
	return __v
}

// AllocateMemoryForImage function as declared in vma/vk_mem_alloc.h:2352
func AllocateMemoryForImage(allocator Allocator, image VkImage, pCreateInfo *AllocationCreateInfo, pAllocation *Allocation, pAllocationInfo *AllocationInfo) VkResult {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cimage, _ := *(*C.VkImage)(unsafe.Pointer(&image)), cgoAllocsUnknown
	cpCreateInfo, _ := pCreateInfo.PassRef()
	cpAllocation, _ := (*C.VmaAllocation)(unsafe.Pointer(pAllocation)), cgoAllocsUnknown
	cpAllocationInfo, _ := pAllocationInfo.PassRef()
	__ret := C.vmaAllocateMemoryForImage(callocator, cimage, cpCreateInfo, cpAllocation, cpAllocationInfo)
	__v := (VkResult)(__ret)
	return __v
}

// FreeMemory function as declared in vma/vk_mem_alloc.h:2360
func FreeMemory(allocator Allocator, allocation Allocation) {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	callocation, _ := *(*C.VmaAllocation)(unsafe.Pointer(&allocation)), cgoAllocsUnknown
	C.vmaFreeMemory(callocator, callocation)
}

// GetAllocationInfo function as declared in vma/vk_mem_alloc.h:2380
func GetAllocationInfo(allocator Allocator, allocation Allocation, pAllocationInfo *AllocationInfo) {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	callocation, _ := *(*C.VmaAllocation)(unsafe.Pointer(&allocation)), cgoAllocsUnknown
	cpAllocationInfo, _ := pAllocationInfo.PassRef()
	C.vmaGetAllocationInfo(callocator, callocation, cpAllocationInfo)
}

// TouchAllocation function as declared in vma/vk_mem_alloc.h:2399
func TouchAllocation(allocator Allocator, allocation Allocation) uint32 {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	callocation, _ := *(*C.VmaAllocation)(unsafe.Pointer(&allocation)), cgoAllocsUnknown
	__ret := C.vmaTouchAllocation(callocator, callocation)
	__v := (uint32)(__ret)
	return __v
}

// SetAllocationUserData function as declared in vma/vk_mem_alloc.h:2416
func SetAllocationUserData(allocator Allocator, allocation Allocation, pUserData unsafe.Pointer) {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	callocation, _ := *(*C.VmaAllocation)(unsafe.Pointer(&allocation)), cgoAllocsUnknown
	cpUserData, _ := pUserData, cgoAllocsUnknown
	C.vmaSetAllocationUserData(callocator, callocation, cpUserData)
}

// CreateLostAllocation function as declared in vma/vk_mem_alloc.h:2431
func CreateLostAllocation(allocator Allocator, pAllocation *Allocation) {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cpAllocation, _ := (*C.VmaAllocation)(unsafe.Pointer(pAllocation)), cgoAllocsUnknown
	C.vmaCreateLostAllocation(callocator, cpAllocation)
}

// MapMemory function as declared in vma/vk_mem_alloc.h:2469
func MapMemory(allocator Allocator, allocation Allocation, ppData *unsafe.Pointer) VkResult {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	callocation, _ := *(*C.VmaAllocation)(unsafe.Pointer(&allocation)), cgoAllocsUnknown
	cppData, _ := ppData, cgoAllocsUnknown
	__ret := C.vmaMapMemory(callocator, callocation, cppData)
	__v := (VkResult)(__ret)
	return __v
}

// UnmapMemory function as declared in vma/vk_mem_alloc.h:2478
func UnmapMemory(allocator Allocator, allocation Allocation) {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	callocation, _ := *(*C.VmaAllocation)(unsafe.Pointer(&allocation)), cgoAllocsUnknown
	C.vmaUnmapMemory(callocator, callocation)
}

// FlushAllocation function as declared in vma/vk_mem_alloc.h:2494
func FlushAllocation(allocator Allocator, allocation Allocation, offset uint, size uint) {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	callocation, _ := *(*C.VmaAllocation)(unsafe.Pointer(&allocation)), cgoAllocsUnknown
	coffset, _ := (C.VkDeviceSize)(offset), cgoAllocsUnknown
	csize, _ := (C.VkDeviceSize)(size), cgoAllocsUnknown
	C.vmaFlushAllocation(callocator, callocation, coffset, csize)
}

// InvalidateAllocation function as declared in vma/vk_mem_alloc.h:2508
func InvalidateAllocation(allocator Allocator, allocation Allocation, offset uint, size uint) {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	callocation, _ := *(*C.VmaAllocation)(unsafe.Pointer(&allocation)), cgoAllocsUnknown
	coffset, _ := (C.VkDeviceSize)(offset), cgoAllocsUnknown
	csize, _ := (C.VkDeviceSize)(size), cgoAllocsUnknown
	C.vmaInvalidateAllocation(callocator, callocation, coffset, csize)
}

// CheckCorruption function as declared in vma/vk_mem_alloc.h:2526
func CheckCorruption(allocator Allocator, memoryTypeBits uint32) VkResult {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cmemoryTypeBits, _ := (C.uint32_t)(memoryTypeBits), cgoAllocsUnknown
	__ret := C.vmaCheckCorruption(callocator, cmemoryTypeBits)
	__v := (VkResult)(__ret)
	return __v
}

// Defragment function as declared in vma/vk_mem_alloc.h:2592
func Defragment(allocator Allocator, pAllocations *Allocation, allocationCount uint, pAllocationsChanged *uint32, pDefragmentationInfo *DefragmentationInfo, pDefragmentationStats *DefragmentationStats) VkResult {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cpAllocations, _ := (*C.VmaAllocation)(unsafe.Pointer(pAllocations)), cgoAllocsUnknown
	callocationCount, _ := (C.size_t)(allocationCount), cgoAllocsUnknown
	cpAllocationsChanged, _ := (*C.VkBool32)(unsafe.Pointer(pAllocationsChanged)), cgoAllocsUnknown
	cpDefragmentationInfo, _ := pDefragmentationInfo.PassRef()
	cpDefragmentationStats, _ := pDefragmentationStats.PassRef()
	__ret := C.vmaDefragment(callocator, cpAllocations, callocationCount, cpAllocationsChanged, cpDefragmentationInfo, cpDefragmentationStats)
	__v := (VkResult)(__ret)
	return __v
}

// BindBufferMemory function as declared in vma/vk_mem_alloc.h:2612
func BindBufferMemory(allocator Allocator, allocation Allocation, buffer VkBuffer) VkResult {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	callocation, _ := *(*C.VmaAllocation)(unsafe.Pointer(&allocation)), cgoAllocsUnknown
	cbuffer, _ := *(*C.VkBuffer)(unsafe.Pointer(&buffer)), cgoAllocsUnknown
	__ret := C.vmaBindBufferMemory(callocator, callocation, cbuffer)
	__v := (VkResult)(__ret)
	return __v
}

// BindImageMemory function as declared in vma/vk_mem_alloc.h:2629
func BindImageMemory(allocator Allocator, allocation Allocation, image VkImage) VkResult {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	callocation, _ := *(*C.VmaAllocation)(unsafe.Pointer(&allocation)), cgoAllocsUnknown
	cimage, _ := *(*C.VkImage)(unsafe.Pointer(&image)), cgoAllocsUnknown
	__ret := C.vmaBindImageMemory(callocator, callocation, cimage)
	__v := (VkResult)(__ret)
	return __v
}

// CreateBuffer function as declared in vma/vk_mem_alloc.h:2660
func CreateBuffer(allocator Allocator, pBufferCreateInfo *VkBufferCreateInfo, pAllocationCreateInfo *AllocationCreateInfo, pBuffer *VkBuffer, pAllocation *Allocation, pAllocationInfo *AllocationInfo) VkResult {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cpBufferCreateInfo, _ := pBufferCreateInfo.PassRef()
	cpAllocationCreateInfo, _ := pAllocationCreateInfo.PassRef()
	cpBuffer, _ := (*C.VkBuffer)(unsafe.Pointer(pBuffer)), cgoAllocsUnknown
	cpAllocation, _ := (*C.VmaAllocation)(unsafe.Pointer(pAllocation)), cgoAllocsUnknown
	cpAllocationInfo, _ := pAllocationInfo.PassRef()
	__ret := C.vmaCreateBuffer(callocator, cpBufferCreateInfo, cpAllocationCreateInfo, cpBuffer, cpAllocation, cpAllocationInfo)
	__v := (VkResult)(__ret)
	return __v
}

// DestroyBuffer function as declared in vma/vk_mem_alloc.h:2679
func DestroyBuffer(allocator Allocator, buffer VkBuffer, allocation Allocation) {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cbuffer, _ := *(*C.VkBuffer)(unsafe.Pointer(&buffer)), cgoAllocsUnknown
	callocation, _ := *(*C.VmaAllocation)(unsafe.Pointer(&allocation)), cgoAllocsUnknown
	C.vmaDestroyBuffer(callocator, cbuffer, callocation)
}

// CreateImage function as declared in vma/vk_mem_alloc.h:2685
func CreateImage(allocator Allocator, pImageCreateInfo *VkImageCreateInfo, pAllocationCreateInfo *AllocationCreateInfo, pImage *VkImage, pAllocation *Allocation, pAllocationInfo *AllocationInfo) VkResult {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cpImageCreateInfo, _ := pImageCreateInfo.PassRef()
	cpAllocationCreateInfo, _ := pAllocationCreateInfo.PassRef()
	cpImage, _ := (*C.VkImage)(unsafe.Pointer(pImage)), cgoAllocsUnknown
	cpAllocation, _ := (*C.VmaAllocation)(unsafe.Pointer(pAllocation)), cgoAllocsUnknown
	cpAllocationInfo, _ := pAllocationInfo.PassRef()
	__ret := C.vmaCreateImage(callocator, cpImageCreateInfo, cpAllocationCreateInfo, cpImage, cpAllocation, cpAllocationInfo)
	__v := (VkResult)(__ret)
	return __v
}

// DestroyImage function as declared in vma/vk_mem_alloc.h:2704
func DestroyImage(allocator Allocator, image VkImage, allocation Allocation) {
	callocator, _ := *(*C.VmaAllocator)(unsafe.Pointer(&allocator)), cgoAllocsUnknown
	cimage, _ := *(*C.VkImage)(unsafe.Pointer(&image)), cgoAllocsUnknown
	callocation, _ := *(*C.VmaAllocation)(unsafe.Pointer(&allocation)), cgoAllocsUnknown
	C.vmaDestroyImage(callocator, cimage, callocation)
}
